package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map // 并发安全的 map

	// 启动多个协程并发写入数据
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i) // 存储键值对
		}(i)
	}

	wg.Wait()

	// 按key读取
	val, ok := m.Load(1)
	if ok {
		println(val)
	}
	// todo 存在则返回，否则写入
	// todo 存在则删除并返回

	// 读取并打印数据
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v -> %v\n", key, value)
		return true
	})
}
