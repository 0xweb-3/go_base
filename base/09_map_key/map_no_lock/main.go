package main

import "fmt"

func main() {
	//var count int64
	//
	//// 原子更新
	//atomic.AddInt64(&count, 1)
	//fmt.Println("Count:", count) // Count: 1
	//
	//// 原子读取
	//value := atomic.LoadInt64(&count)
	//fmt.Println("Value:", value) // Value: 1

	// 如果不希望使用加锁，且要保证多个 goroutine 协作更新 map，可以通过一个协作机制，比如 channel，来同步更新操作。
	//通过将 map 的操作通过 channel 发送到一个专门的 goroutine 中执行，这样可以避免多个 goroutine 直接操作 map，从而避免并发问题。

	type MapUpdate struct {
		Key   string
		Value int
	}

	m := make(map[string]int)
	ch := make(chan MapUpdate)

	// 启动一个 goroutine 来处理更新
	go func() {
		for update := range ch {
			m[update.Key] = update.Value
		}
	}()

	// 向 channel 发送更新
	ch <- MapUpdate{"a", 10}
	ch <- MapUpdate{"b", 20}

	// 等待处理
	fmt.Println("Map after updates:", m)
}
