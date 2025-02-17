package main

import (
	"fmt"
	"net/http"
)

//func main() {
//	// 1. 创建http server对象
//	mux := http.NewServeMux()
//
//	// 2. 注册路由和中间件
//	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
//		fmt.Fprint(response, "Hello World!")
//	})
//
//	// 3. 启动http服务，传入监听地址和多路复用器
//	err := http.ListenAndServe(":8080", mux)
//	if err != nil {
//		println(err)
//	}
//}

func main() {
	// 1. 创建http server对象
	//mux := http.NewServeMux()

	// 2. 注册路由和中间件
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Hello World!")
	})

	// 3. 启动http服务，传入监听地址和多路复用器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err)
	}
}
