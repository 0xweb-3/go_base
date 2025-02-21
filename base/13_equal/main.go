package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var b1 []byte
	b2 := []byte{} // 开辟了内存空间

	fmt.Println(reflect.DeepEqual(b1, b2)) // false
	fmt.Println(bytes.Equal(b1, b2))       // true

	var b3 []byte
	var b4 []byte

	fmt.Println(reflect.DeepEqual(b3, b4)) // true
	fmt.Println(bytes.Equal(b3, b4))       // true

	b5 := []byte{}
	b6 := []byte{}

	fmt.Println(reflect.DeepEqual(b5, b6)) // true
	fmt.Println(bytes.Equal(b5, b6))       // true
}
