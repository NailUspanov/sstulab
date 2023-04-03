package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

func main() {

	google, _ := url.Parse("https://www.google.com")
	bing, _ := url.Parse("https://www.bing.com")
	duckduck, _ := url.Parse("https://www.duckduckgo.com")
	pool := ServerPool{
		backends: []*Backend{
			{
				URL:          google,
				Alive:        true,
				mux:          sync.RWMutex{},
				ReverseProxy: httputil.NewSingleHostReverseProxy(google),
			},
			{
				URL:          bing,
				Alive:        true,
				mux:          sync.RWMutex{},
				ReverseProxy: httputil.NewSingleHostReverseProxy(bing),
			},
			{
				URL:          duckduck,
				Alive:        true,
				mux:          sync.RWMutex{},
				ReverseProxy: httputil.NewSingleHostReverseProxy(duckduck),
			},
		},
		current: 0,
	}
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		pool.lb(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func (s *ServerPool) lb(w http.ResponseWriter, r *http.Request) {
	peer := s.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

func (s *ServerPool) GetNextPeer() *Backend {
	// loop entire backends to find out an Alive backend
	next := s.NextIndex()
	l := len(s.backends) + next // start from next and move a full cycle
	for i := next; i < l; i++ {
		idx := i % len(s.backends) // take an index by modding with length
		// if we have an alive backend, use it and store if its not the original one
		if s.backends[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx)) // mark the current one
			}
			return s.backends[idx]
		}
	}
	return nil
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}
