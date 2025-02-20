package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 协程对同一块内存空间的竞争导致的
func counter1() {
	cnt := int(0)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt += 1
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

// 使用加锁方式
func counter2() {
	cnt := int(0)
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			cnt += 1
			defer func() {
				lock.Unlock()
				wg.Done()
			}()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

// 使用channel
func counter3() {
	cnt := int(0)
	wg := sync.WaitGroup{}
	// 如果缓冲区大小设置为更大的值（例如 2 或更多），
	//则会允许更多的 goroutine 同时进入临界区进行操作，
	//这可能会引起竞态条件，导致结果不可靠，特别是当多个
	//goroutine 同时修改 cnt 时，可能会发生数据不一致的情况。
	//
	//如果缓冲区设置为 0（无缓冲的 channel），
	//则会强制执行同步操作，但可能会导致更多的上下文切换开销，
	//因为每次 goroutine 执行时，都需要在发送和接收时等待对方。
	ch := make(chan struct{}, 1)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			ch <- struct{}{}
			defer func() {
				<-ch
				wg.Done()
			}()
			cnt += 1
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

func counter4() {
	cnt := int64(0)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&cnt, 1)
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

func main() {
	//counter1() //9613
	//counter2()

	counter3()
}
