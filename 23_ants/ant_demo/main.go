package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

// 启动一个定时器
func Ticker(f func(), d time.Duration) {
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("当前时间：", t)
				go f()
			}
		}
	}()
}

func init() {
	Ticker(func() {
		log.Println("current go routine num:", runtime.NumGoroutine())
	}, 2*time.Second)
}

func main() {
	// 创建一个ants池，最多允许5个goroutine同时执行
	p, _ := ants.NewPool(5)
	defer p.Release()

	// 定义一个计数器，用于统计处理的任务数
	var counter int64

	for i := 0; i < 100; i++ {
		x := i
		err := p.Submit(func() {
			// 模拟任务处理耗时
			time.Sleep(time.Second)
			log.Println("执行任务", x)
			// 增加计数器
			atomic.AddInt64(&counter, 1)
		})
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("等待执行的任务数量：", p.Waiting())

	// 当运行任务数为0时，说明任务执行完毕
	for p.Running() > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	// 输出计数器的值
	log.Println("counter:", counter)

	//select {}
}
