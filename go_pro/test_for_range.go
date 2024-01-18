package main

import "fmt"

// 数组
func c1() {
	var a = [...]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

// 切片
func c2() {
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

// map
func c3() {
	// key : value
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "ghz@gmail.com"
	for key, value := range m {
		fmt.Printf("%v:%v", key, value)
	}
}

// 字符串
func c4() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}

func main() {
	//c1()
	//c2()
	//c3()
	c4()
}
