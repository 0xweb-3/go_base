package main

import "sync"

func task(i int, wg *sync.WaitGroup) {
	println(i)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		x := i
		go func() {
			task(x, &wg) // wg必须传递引用类型
		}()
	}
	wg.Wait() // 等待wg计数器中的值变为0
}
