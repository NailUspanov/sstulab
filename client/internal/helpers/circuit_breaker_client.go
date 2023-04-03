package helpers

import (
	"net/http"
)

//func NewCircuitBreakerClient(endpoint string, cb *CircuitBreaker) (*http.Client, error) {
//	transport := http.DefaultTransport.(*http.Transport).Clone()
//
//	// Create a RoundTripper that uses the circuit breaker
//	transport.RegisterProtocol("http", &circuitBreakerRoundTripper{
//		cb:       cb,
//		delegate: http.DefaultTransport.(*http.Transport),
//	})
//
//	client := &http.Client{
//		Timeout:   10 * time.Second,
//		Transport: transport,
//	}
//
//	fmt.Println("FUCK YOU!")
//
//	// Check that the endpoint is valid by sending a HEAD request
//	_, err := client.Head(endpoint)
//	if err != nil {
//		fmt.Println("FUCK YOU! TOOOOO")
//		return nil, err
//	}
//	fmt.Println("OIMG!")
//
//	return client, nil
//}

type circuitBreakerRoundTripper struct {
	cb       *CircuitBreaker
	delegate http.RoundTripper
}

//func (rt *circuitBreakerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
//	err := rt.cb.Execute()
//	if err != nil {
//		return nil, err
//	}
//
//	return rt.delegate.RoundTrip(req)
//}
