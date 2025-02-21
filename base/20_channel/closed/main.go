package main

import (
	"fmt"
	"time"
)

func f1() {
	ch := make(chan int)
	close(ch)

	go func() {
		println("开始接收数据")
		fmt.Println(<-ch)
		println("接收数据完毕")
	}()

	go func() {
		println("开始发送数据")
		ch <- 1
		println("发送数据完毕")
	}()

	time.Sleep(1 * time.Second)
}

func f2() {
	ch := make(chan int)
	close(ch)
	for {
		select {
		case c := <-ch:
			fmt.Println("读取<-ch", c)
			time.Sleep(time.Second)
		}
	}

}

func f3() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}

		//close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}

func f4() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()

	// 判断channel关闭的方式
	for {
		select {
		case x, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}

			fmt.Println(x)
		}
	}
}

func main() {
	// 不能向已经关闭的channel发送数据
	//f1() // panic: send on closed channel

	//f2() // 可以从关闭的channel中读取数据，并不会阻塞

	//f3()

	f4() // channel关闭的判断
}
