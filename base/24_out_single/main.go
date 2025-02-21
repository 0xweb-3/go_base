package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 设置一个路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 创建 HTTP 服务
	server := &http.Server{
		Addr:    ":8080", // 监听端口
		Handler: r,       // 设置 Gin 路由处理器
	}

	// 创建一个信号通道，用于接收系统信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill) // 捕获终止信号

	// 启动 HTTP 服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("服务启动成功，等待停止信号...")

	// 等待接收终止信号
	<-sigChan
	fmt.Println("收到停止信号，正在优雅关闭服务...")

	// 创建一个 5 秒的超时时间来等待正在处理的请求完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务关闭失败：%s", err)
	}

	fmt.Println("服务已成功关闭")
}
