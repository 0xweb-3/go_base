package main

import "fmt"

func f(x int) int {
	defer func() {
		x += x // 修改返回值
	}()

	return x + x
}

func f2() {
	x := 0
	defer func(paramx int) {
		fmt.Println(paramx) // 0
	}(x)
	x = 1
	fmt.Println("done")
}

func main() {
	f2()
}
