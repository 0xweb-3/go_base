package main

import (
	"errors"
	"fmt"
)

var errNotFount = errors.New("not found")

// 模拟一个函数，返回一个错误，并附带额外的上下文信息
func getResource(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid id: %w", errNotFount)
	}
	return fmt.Sprintf("%d", id), nil
}

func main() {

	// Go 1.13 引入了错误包装（fmt.Errorf）和 errors.Is 和
	//errors.As，可以将错误附加更多上下文信息。
	//这种方法使得错误信息更加丰富，易于追踪和诊断。
	_, err := getResource(-1)
	if err != nil {
		if errors.Is(err, errNotFount) {
			fmt.Println("not found")
		} else {
			fmt.Println(err.Error())
		}
	}
}
