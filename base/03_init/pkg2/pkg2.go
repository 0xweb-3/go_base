package pkg2

import (
	"fmt"
	"github.com/0xweb-3/go_base/03_init/pkg3"
)

// 不同package且相互依赖，最后被依赖的最先被执行
var Pkg2 = pkg3.Pkg3

func init() {
	fmt.Println("pkg2-pkg2")
}
