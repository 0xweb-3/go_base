package main

import "fmt"

func main() {
	//// 1.使用 sync.Mutex 或 sync.RWMutex
	//var m = make(map[int]int)
	//var mu sync.Mutex
	//
	////写操作，添加键值对
	//mu.Lock()
	//m[1] = 100
	//mu.Unlock()
	//
	//// 读操作，获取键值对
	//mu.Lock()
	//val := m[1]
	//mu.Unlock()
	//fmt.Println("val:", val)
	//
	////如果读操作远多于写操作，可以使用 sync.RWMutex 来优化性能。sync.RWMutex 允许多个读操作并发执行，但写操作仍然是独占的。
	//var m2 = make(map[int]int)
	//var rwmu sync.RWMutex
	//
	//// 写操作，添加键值对
	//rwmu.Lock()
	//m2[1] = 100
	//rwmu.Unlock()
	//
	//// 读操作，获取键值对
	//rwmu.RLock()
	//value := m2[1]
	//rwmu.RUnlock()
	//fmt.Println(value)

	// 通道 ch 用作互斥量，确保同一时间只有一个 goroutine 可以操作 map。
	m := make(map[int]int)
	ch := make(chan struct{}, 1)

	//写操作
	ch <- struct{}{} // 进入临界区
	m[1] = 100
	<-ch // 离开临界区

	// 读操作
	ch <- struct{}{}
	value := m[1]
	<-ch
	fmt.Println(value)
}
