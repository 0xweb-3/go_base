package main

import "fmt"

func change1(s []string) {
	// 传递进来的是地址，直接复制的地址
	fmt.Printf("change1 s pointer%p\n", s)

	// 此时将s赋值新的地址
	s = []string{"a", "b", "c"}
	fmt.Printf("changed s pointer%p\n", s) // 和前面打印的地址不一致
}

func change2(s []string) {
	// V只是一个值拷贝的副本，这里直接修改值不会修改原来的值
	for i, v := range s {
		v = "B"
		fmt.Println(i, v)
	}
}

func change3(s []string) {
	s[1] = "C"
}

type person struct {
	name string
	age  int
}

type persons struct {
	names []string
}

func changeStruct1(p person) {
	p.name = "xin"
}

func changeStruct2(ps persons) {
	fmt.Printf("changeStruct2 ps pointer%p\n", ps.names)
	ps.names[1] = "xin"
	fmt.Printf("changeStruct2 ed  ps pointer%p\n", ps.names)
}

func main() {
	//a := []string{"x", "y", "z"}
	//change1(a)
	//fmt.Printf("a pointer%p\n", a)

	//change2(a)
	//fmt.Println(a)

	//change3(a)
	//fmt.Println(a) // [x C z]

	//p := person{
	//	name: "xxx",
	//	age:  10,
	//}
	//changeStruct1(p)
	//fmt.Println(p) // {xxx 10}

	ps := persons{
		names: []string{"a", "b", "c"},
	}

	changeStruct2(ps)
	fmt.Println(ps)
	fmt.Printf("ps pointer%p\n", ps.names)
}
