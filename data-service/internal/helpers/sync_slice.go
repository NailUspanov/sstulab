package helpers

import "sync"

type ThreadSafeSlice struct {
	sync.Mutex
	slice []interface{}
}

func (t *ThreadSafeSlice) Append(item interface{}) {
	t.Lock()
	defer t.Unlock()
	t.slice = append(t.slice, item)
}

func (t *ThreadSafeSlice) Get(index int) interface{} {
	t.Lock()
	defer t.Unlock()
	return t.slice[index]
}

func (t *ThreadSafeSlice) Len() int {
	t.Lock()
	defer t.Unlock()
	return len(t.slice)
}

func (t *ThreadSafeSlice) Delete(index int) {
	t.Lock()
	defer t.Unlock()
	t.slice = append(t.slice[:index], t.slice[index+1:]...)
}

func (t *ThreadSafeSlice) GetSlice() []interface{} {
	return t.slice
}
