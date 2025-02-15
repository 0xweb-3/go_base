package main

import (
	"os"
	"strconv"
	"sync"
)

func processFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 文件的处理
	return nil
}

func processFile2() error {
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			processFile("test" + strconv.Itoa(i) + ".txt")
		}(i)
	}

	wg.Wait()
	return nil
}

func main() {
	for i := 0; i < 1000000; i++ {
		processFile("test" + strconv.Itoa(i) + ".txt")
	}
	return
}
