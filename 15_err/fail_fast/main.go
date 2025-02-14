package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	// 假设出错
	return errors.New("something went wrong")
}

func main() {
	// 直接检查错误并返回
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 继续执行其他逻辑
}
