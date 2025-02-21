package main

import (
	"fmt"
	"log"
)

// 统一处理错误的函数
func handleError(err error) {
	if err != nil {
		log.Fatal(err) // 或者 log.Println(err)，根据需求决定
	}
}

func main() {
	// 示例错误
	err := fmt.Errorf("an example error")
	handleError(err)
}
