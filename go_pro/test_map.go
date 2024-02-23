package main

import "fmt"

//map是一种key:value键值对的数据结构容器。map内部实现是哈希表(hash).map最重要的一点是通过key来快速检索数据。key类似于索引，指向数据的值.map是引用类型的。

func m1() {
	var m1 map[string]string     // 1 声明变量，默认map是nil
	m1 = make(map[string]string) //2 使用make函数
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

// 初始化
func m2() {
	//map是无序的
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	fmt.Printf("m1: %v\n", m1)
	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["email"] = "tom@gmail.com"
	fmt.Printf("m2: %v\n", m2)
}

func m3() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}

// 判断某个键是否存在
func m4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v： %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("-------------")
	v, ok = m1[k2]
	fmt.Printf("v： %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}

func main() {
	//m1()
	//m2()
	//m3()
	m4()
}
