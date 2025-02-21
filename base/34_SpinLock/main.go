package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	lock int32
}

// Lock 方法实现自旋锁
func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.lock, 0, 1) {
		// 自旋，持续检查 lock 是否为 0，如果是 0 就尝试获取锁
	}
}

// Unlock 方法释放自旋锁
func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.lock, 0)
}

func main() {
	var spinLock SpinLock
	var wg sync.WaitGroup

	// 启动多个 goroutine 并发执行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 获取锁
			spinLock.Lock()
			defer spinLock.Unlock()

			// 模拟一些工作
			fmt.Printf("goroutine %d is working\n", i)
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("goroutine %d finished\n", i)
		}(i)
	}

	wg.Wait()
}
