package main

import "fmt"

// 定义一个接口
type Marker interface {
	Mark() string
}

// 定义一个空结构体，作为接口实现的占位符
type Empty struct{}

func (e *Empty) Mark() string {
	return "This is a marker!"
}

func main() {
	var m Marker = &Empty{}
	fmt.Println(m.Mark()) // 输出: This is a marker!
}
