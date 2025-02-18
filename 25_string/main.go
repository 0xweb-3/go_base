package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := "abc"
	s2 := "中文"

	// len获得的是字节数，一个中文两个字节。
	fmt.Println(len(s1), unsafe.Sizeof(s1)) // 3 16
	fmt.Println(len(s2), unsafe.Sizeof(s2)) // 6 16

	r := []rune(s2)
	fmt.Println(len(r), unsafe.Sizeof(r)) // 2 24
}
