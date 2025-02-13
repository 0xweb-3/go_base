package main

import "fmt"

func main() {
	//数组的零值和切片的零值不同
	//一个数组的零值是数组的所有元素均为对应数组元素类型的零值
	//arr1 := [10]int{}
	//fmt.Println(arr1) // [0 0 0 0 0 0 0 0 0 0]
	//
	//var arr2 [10]int
	//fmt.Println(arr2) // [0 0 0 0 0 0 0 0 0 0]

	// 有零值
	//arr3 := new([10]int)
	//fmt.Println(arr3) // &[0 0 0 0 0 0 0 0 0 0]
	//var arr4 *[10]int = &[10]int{}
	//fmt.Println(arr4) // &[0 0 0 0 0 0 0 0 0 0]

	// 为nil,无零值，直接使用会报错
	//var arr5 *[10]int
	//fmt.Println(arr5)
	//fmt.Println(arr5 == arr3) //指向的指针地址是不同的

	var s1 = []int{}
	fmt.Println(s1) // []
	var s2 = new([]int)
	fmt.Println(s2)        // &[]
	fmt.Println(s2 == nil) //false

	s1 = append(s1, 1)
	*s2 = append(*s2, 1)
	fmt.Println(s1) // [1]
	fmt.Println(s2) // &[1]

	var s3 = make([]int, 1)
	fmt.Println(s3) // [0]

	var s4 []int
	fmt.Println(s4)        // []
	fmt.Println(s4 == nil) // true

	var s5 *[]int

	tmp := []int{}
	s5 = &tmp
	tmp1 := append(tmp, 1)
	s5 = &tmp1
	fmt.Println(s5) // &[1]

	var x = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(x) // [0 77 88 0 44 55 66]

}
