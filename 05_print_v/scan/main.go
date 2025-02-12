package main

import "fmt"

func main() {
	//Scan
	var (
		appName string
		version int
	)
	//fmt.Println("请输入name:")
	//fmt.Scan(&appName)
	//fmt.Println("请输入version")
	//fmt.Scan(&version)
	//fmt.Printf("fmt.Scan appName:%s version:%d \n", appName, version)

	// Scanf
	//_, err := fmt.Scanf("name= %s ver= %s", &appName, &version)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("fmt.Scanf appName:%s version:%s \n", appName, version)

	// ScanIn
	//fmt.Println("请输入name")
	//fmt.Scanln(&appName) // 遇到回车结束
	//fmt.Println("请输入version")
	//fmt.Scanln(&version)
	//fmt.Printf("fmt.
	//Scan appName:%s version:%s \n", appName, version)

	//reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象 //fmt.PrintIn("请输入：")
	//text, _ := reader.ReadString('\n')  // 读到换行
	////text, _ := reader.ReadString(' ')   // 读到换行
	//fmt.Println(text)

	// go run -Idflags "-X 'main.appName=test' -X 'main.version=1'" main.go
	fmt.Printf("appName:%s version:%s \n", appName, version)
}
