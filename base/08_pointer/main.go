package main

import "fmt"

func f1(x int) {
	x = 10
}

func f2(x *int) {
	// 实参x的副本跟c指向同一个地址，所以修改副本指向的值也会影响到函数外c的值
	*x = 10
	// 这里直接修改了实参x的副本地址，相当于让实参x指向了一个新的地址，所以这里改变x的地址并不会影响函数外c的值
	x = nil
}

func main() {
	x := 0
	f1(x)
	fmt.Println(x)
	f2(&x)
	fmt.Println(x)
}
