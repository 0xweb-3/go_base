package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func slowOperation(ctx context.Context) (string, error) {
	// 获取request_id
	fmt.Println("get request_id from ctx", ctx.Value("request_id"))

	//使用一个 select 语句来监听 context 的状态变化
	select {
	case <-time.After(time.Second * 3):
		return "success", nil
	case <-ctx.Done(): // 如果 context 被取消或者超时，返回相应的错误信息
		return "", ctx.Err()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestId := "1111"

	ctr := context.WithValue(r.Context(), "request_id", requestId)

	ctx, cancel := context.WithTimeout(ctr, 1*time.Second)
	defer cancel()

	//另一种超时设置方式
	//deadline := time.Now).Add(1 * time.Second)
	//ctx, cancel := context.WithDeadline(ctr, deadline)

	result, err := slowOperation(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}

	// 如果操作成功，返回结果
	fmt.Fprintln(w, result)

}

func main() {
	//创建一个 HTTP 服务器，并注册处理函数
	http.HandleFunc("/", handler)
	fmt.Println("start http server on :9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println()
	}
}
