package main

import "fmt"

func fn() {
	println("call fn")
	f1()
	println("exit fn")
}

func f1() {
	println("call f1")
	defer func() {
		fmt.Println("defer before painc in f1")
	}()
	panic("panic in f1")
	defer func() {
		fmt.Println("defer after painc in f1")
	}()
	f2()
	panic("exit f1")
}

func f2() {
	println("call f2")
	println("call fn2")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到一个painc", err)
			fmt.Println("防止程序崩溃")
		}
	}()

	println("call main")
	fn()
	println("exit main")
}
