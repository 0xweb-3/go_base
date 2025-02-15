package main

import "fmt"

func main() {
	var c chan int

	var c2 chan<- int // 单向只写
	var c3 <-chan int // 单向只读

	c2 = c
	c3 = c

	fmt.Println(c2, c3)

	//readOnlyChanWithBuff := make(<-chan int, 2)  // 单向只读channel，并带缓冲
	//readOnlyChan := make(<-chan int)             // 单向只读channel，不带缓冲
	//writeOnlyChanWithBuff := make(chan<- int, 4) // 单向只写channel，并带缓冲区
}
