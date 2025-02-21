package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("hello world")
			time.Sleep(time.Second)
		}
	}()

	select {}
}
