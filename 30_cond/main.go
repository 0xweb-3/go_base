package main

import (
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)

	go func() {
		//Wait0）方法的调用必须加锁
		cond.L.Lock()
		println("wait1 start")
		// 开始等待信号
		cond.Wait()
		println("wait1 end....")
		cond.L.Unlock()
	}()
	go func() {
		//Wait0）方法的调用必须加锁
		cond.L.Lock()
		println("wait2 start")
		// 开始等待信号
		cond.Wait()
		println("wait2 end....")
		cond.L.Unlock()
	}()

	time.Sleep(3 * time.Second)
	//调用Signal方法只能让上面第一个Wait()得到释放
	cond.Signal()
	// 调用Broadcast()方法能释放全部WaitO
	cond.Broadcast()
	time.Sleep(time.Second * 10)
}
