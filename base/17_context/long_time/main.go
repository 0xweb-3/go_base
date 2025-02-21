package main

import (
	"context"
	"time"
)

func longRunningFunction(ctx context.Context, arg1 string, arg2 int) error {
	select {
	case <-time.After(5 * time.Second):
	//执行操作
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, time.Second*4)
	defer cancel()
	if err := longRunningFunction(ctx, "", 1); err != nil {
		// 如果发生错误，调用cancel()来取消请求
		cancel()
	}
}
