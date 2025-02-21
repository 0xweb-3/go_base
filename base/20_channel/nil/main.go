package main

import (
	"fmt"
	"time"
)

func ChannelNil() {
	var ch chan struct{}
	ch2 := make(chan struct{})
	println("ch == nil", ch == nil)   // true
	println("ch2 == nil", ch2 == nil) // false

	close(ch2)
	// 关闭后的channel不为nil
	println("关闭后的channel：ch2 == nil", ch2 == nil) // false
}

// fatal error: all goroutines are asleep - deadlock!
func f1() {
	var c chan int // 声明了一个 nil channel
	go func() {
		c <- 1 // 尝试向 nil channel 发送数据，会阻塞
	}()

	println(<-c) // 尝试从 nil channel 接收数据，会阻塞
}

// 阻塞到channel退出
func f2() {
	var c chan int
	go func() {
		println("开始发送")
		c <- 1
		fmt.Println("阻塞在发送中") // 不会执行
	}()

	go func() {
		println("开始接收")
		println(<-c)
		fmt.Println("阻塞在接收中") // 不会执行
	}()
	time.Sleep(time.Second)
}

func closeNilChanel() {
	var ch chan struct{}

	close(ch)
}

func main() {
	//ChannelNil()
	//f1()
	//f2()

	//panic: close of nil channel
	closeNilChanel()
}
