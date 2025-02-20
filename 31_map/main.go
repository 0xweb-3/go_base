package main

import (
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[any]any
}

func (sm *SafeMap) Get(key any) any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.m[key]
}

func (sm *SafeMap) Set(key, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Delete(key any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

func (sm *SafeMap) Range(f func(key, value any) bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for k, v := range sm.m {
		if !f(k, v) {
			break
		}
	}
}

func main() {
	m := SafeMap{
		mu: sync.RWMutex{},
		m:  map[any]any{},
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//x := i
		go func(index int) {
			m.Set(index, index)
			wg.Done()
		}(i) // 进行了一次值传递
	}

	wg.Wait()
	m.Range(func(key, value any) bool {
		println(key, value)
		return true
	})
}
