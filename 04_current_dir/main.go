package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("./")
	//相对路径指的是相对命令执行的位置，不是执行代码的位置
	fmt.Println(path, err)
}
