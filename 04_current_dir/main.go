package main

import "github.com/0xweb-3/go_base/04_current_dir/root_dir/config"

func main() {
	//path, err := filepath.Abs("./")
	////相对路径指的是相对命令执行的位置，不是执行代码的位置
	//fmt.Println(path, err)

	//config.GetWorkPath()
	//config.GetWorkByArg()
	config.GetWorkByExec()
}
