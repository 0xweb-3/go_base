package main

import (
	"fmt"
	"log"
	"time"
)

// 通过go关键字启动
func nativeGo() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}

func nativeGoFunc() {
	for i := 0; i < 10; i++ {
		x := i
		go func() {
			//log.Println(i) // 全是10
			log.Println(x)
		}()
	}
}

type S struct {
	F1 string
	F2 int
	F3 int
}

func goWithPointerData() {
	s1 := &S{
		F1: "xin",
		F2: 1,
		F3: 0,
	}
	for i := 0; i < 10; i++ {
		s2 := s1
		s2.F3 = i
		writeData(s2)
	}
}

func writeData(s *S) {
	//'s是指针类型，在协程外有修改，内部访问的值不确定
	go func() {
		log.Println("F3", s.F3)
	}()
}

func main() {
	//nativeGo()
	//nativeGoFunc()

	goWithPointerData() // 打印值不确定

	time.Sleep(time.Second)
}
