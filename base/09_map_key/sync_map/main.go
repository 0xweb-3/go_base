package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map // // 声明一个sync.Map类型的变量

	// 存储一个键值对
	m.Store("name", "xin")
	fmt.Println(m)

	// 加载一个键对应的值
	if v, ok := m.Load("name"); ok {
		// 如果该键存在，则输出值
		fmt.Println("name:", v)
	}

	//删除一个键值对
	m.Delete("name")

	// 遍历map中的所有键值对
	m.Store("age", 30)         // 存储一个新的键值对
	m.Store("location", "NYC") // 存储一个新的键值对

	m.Range(func(key, value interface{}) bool {
		// 遍历每一个键值对并输出
		fmt.Println(key, ":", value)
		return true // 返回true继续遍历
	})
}
