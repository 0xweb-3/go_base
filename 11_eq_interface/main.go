package main

func main() {
	var a int
	var b interface{} = a
	var c int32
	var d interface{} = c

	println("b=", b, "d=", d)
	println("b == d?", b == d)

	var e error
	var emptyl interface{} // 空接口类型
	println("e =nil ?", e == nil)
	println("emptyl =nil ?", emptyl == nil)
	println("e == emptyl?", e == emptyl)
	println("e=", e, "emptyl=", emptyl)
}
