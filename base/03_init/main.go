package main

import (
	"fmt"
	"github.com/0xweb-3/go_base/03_init/pkg"
	
	"github.com/0xweb-3/go_base/03_init/pkg1"
	"github.com/0xweb-3/go_base/03_init/pkg2"
)

// 同一文件中的init函数，按照代码顺序执行
func init() {
	fmt.Println("main1")
}

func init() {
	fmt.Println("main2")
}

func main() {
	// 同一package中的，按照文件名的顺序执行
	// 被引入的包会先执行
	fmt.Println(pkg.Pkg)

	// 不同package且不相互依赖，按照import顺序执行
	fmt.Println(pkg1.Pkg1, pkg2.Pkg2)
}
