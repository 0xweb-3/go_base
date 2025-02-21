package main

import (
	"fmt"
	"sync"
)

func main() {
	const num = 100

	s := make([]int, num)
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		x := i
		wg.Add(1)
		go func() {
			s[x] = x
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s)
}
