package main

import "fmt"

type Person struct {
	uncompare [0]func() // 定义了一个长度为 0 的数组，数组的元素类型是一个函数类型（func()。不占用内存空间
	Name      string
}

type Person2 struct {
	Name string
	Age  int
}

type Person3 struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "xin",
	}
	fmt.Println(p1)

	//p2 := Person2{
	//	Name: "xin",
	//	Age:  12,
	//}
	//
	//p3 := Person3{
	//	Name: "xin",
	//	Age:  12,
	//}
	//fmt.Println("p2 == p3", p2 == p3) // 不同类型不能比较

	//p5 := struct {
	//	Name 25_string
	//	Age  int
	//}{}
	//
	//p6 := struct {
	//	Name 25_string
	//	Age  int
	//}{}
	//
	//fmt.Println("p5 == p6", p5 == p6) // true  字段顺序需要相同

	//p7 := struct {
	//	uncompare [0]func()
	//	Age       int
	//}{}
	//
	//p8 := struct {
	//	uncompare [0]func()
	//	Age       int
	//}{}
	//fmt.Println("p7 == p8", p7 == p8) // invalid operation: p7 == p8 (struct containing [0]func() cannot be compared)
}
