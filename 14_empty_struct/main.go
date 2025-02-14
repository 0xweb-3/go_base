package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 零占用空间
	st := struct{}{}
	fmt.Println("unsafe.Sizeof", unsafe.Sizeof(st)) // 0

}
