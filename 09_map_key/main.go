package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func structToMap(s interface{}) map[string]interface{} {
	// 获取结构体的值和类型
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// 如果传入的是指针类型，获取它指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	// 创建一个 map 用于存储结果
	result := make(map[string]interface{})

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		result[fieldName] = field.Interface()
	}

	return result
}

func main() {
	//a1 := [3]int{}
	//a2 := [3]int{}
	//fmt.Println(a1 == a2) // 数组是可以比较的
	//
	////包含切片（不可比较类型）的数组是不能比较的
	//a3 := [3][]int{}
	//a4 := [3][]int{}
	//fmt.Println(a3 == a4)

	//m := map[float64]int{}
	//m[1.1] = 1
	//m[1.2] = 2
	//m[0.3] = 5
	//m[0.300000000000000000001] = 1
	//// 在浮点数作为key时，会使用math.Float64bits(key),导致精度丢失
	//fmt.Println(m) // map[0.3:1 1.1:1 1.2:2]

	// new分配的零值的指针是不同的
	//p1 := new(int)
	//p2 := new(int)
	//
	//fmt.Println(p1 == p2) //false
	//
	//m1 := map[*int]int64{}
	//m1[p1] = 1
	//m1[p2] = 2
	//
	//fmt.Println(m1)

	// map转换为struct
	m := map[string]interface{}{
		"Name": "xin",
		"Age":  18,
	}
	var p Person
	if name, ok := m["Name"]; ok {
		p.Name = name.(string)
	}

	if age, ok := m["Age"]; ok {
		p.Age = age.(int)
	}
	fmt.Println(p)

	// struct 转map
	p1 := Person{Name: "John", Age: 30}
	// 将 struct 转换为 map
	m1 := structToMap(p1)
	// 输出结果
	fmt.Println(m1)

	// 使用 json 序列化与反序列化进行转换
	p3 := Person{Name: "John", Age: 30}
	// 将结构体转换为 JSON 字符串
	jsonData, _ := json.Marshal(p3)
	// 将 JSON 字符串转换为 map
	var m3 map[string]interface{}
	json.Unmarshal(jsonData, &m3)
	// 输出结果
	fmt.Println(m3) // 输出: map[age:30 name:John]
	// 将 map 转换为 struct
	var p4 Person
	jsonData2, _ := json.Marshal(m3)
	json.Unmarshal(jsonData2, &p4)

	// 输出结构体
	fmt.Println(p4) // 输出: {John 30}
}
