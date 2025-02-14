package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	m := make(map[int]*int)

	//for key, value := range slice {
	//	// 使用新的变量保存当前值的地址
	//	//val := v
	//
	//	// 在这里 v 是切片元素的值的拷贝
	//	m[key] = &value // 存储的是 v 的地址，v 在每次迭代中都会改变
	//}

	// 优化写法
	for key, _ := range slice {
		m[key] = &slice[key]
	}

	// 输出 map 中保存的地址指向的值
	for k, v := range m {
		fmt.Println(k, "->", *v) // *v 解引用，打印 v 所指向的值
	}
}
