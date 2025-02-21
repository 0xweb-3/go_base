package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type myJob struct {
	name string
}

func (j myJob) Run() {
	fmt.Printf("%s is running\n", j.name, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	// 设置时区的两种方式
	// 1. 通过加载环境变量或者操作操作系统或者go内置的包或文件中加载指定时区
	//loc, err := time.LoadLocation("America/Los_Angeles")
	//if err != nil {
	//	println("load location error", err)
	//	return
	//}
	// 2. 直接创建时区
	loc := time.FixedZone("CST", 8*60*60)
	fmt.Println(time.Now().In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().In(time.UTC))

	// 注册周期性任务的两种方式

	//  创建一个cron对象，WithLocation设置时区，WithSeconds秒级别支持
	c := cron.New(cron.WithLocation(loc), cron.WithSeconds())

	// 添加每分钟执行一次的任务
	id1, err := c.AddFunc("* * * * * ", func() {
		// 执行的任务
		println("每分钟执行任务。。。。。。", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		println("id1 err:", err)
	}
	// 添加每秒钟执行一次的任务
	id2, err := c.AddFunc("* * * * * *", func() {
		// 执行的任务
		println("每秒钟执行任务。。。。。。", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		println("id2 err:", err)
	}

	job := myJob{}
	job.name = "id3"
	// 添加每秒钟执行的任务
	id3, err := c.AddJob("* * * * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(job))
	if err != nil {
		println("id3 err:", err)
	}

	// 启动定时任务
	c.Start()

	// 查看定时任务
	for _, e := range c.Entries() {
		fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v \n", e.ID, e.Next, e.Prev, e.Schedule)
	}

	//查看指定的定时任务
	e := c.Entry(id1)
	fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v\n", e.ID, e.Next, e.Prev, e.Schedule)
	e = c.Entry(id2)
	fmt.Printf("ID: %d, Next: %v, Prev: %v, Schedule: %v\n", e.ID, e.Next, e.Prev, e.Schedule)

	// 删除指定的定时任务
	c.Remove(id3)

	// 停止定时任务
	defer c.Stop()

	// 主程序阻塞
	select {}
}
