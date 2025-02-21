package err_struct

import "fmt"

// 定义自定义错误类型
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code:%d message:%s", e.Code, e.Message)
}

// 模拟一个可能出错的函数
func doSomething() error {
	return &MyError{
		Code:    404,
		Message: "Item not found",
	}
}

func main() {

}
