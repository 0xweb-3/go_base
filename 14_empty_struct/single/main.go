package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		// 完成一些任务
		fmt.Println("Worker is doing some work...")
		time.Sleep(3 * time.Second)
		// 发送空结构体作为信号
		ch <- struct{}{}
	}()

	// 等待信号
	<-ch
	fmt.Println("Received signal, worker has finished.")
}
