package main

import "fmt"

//不能放外面
//email := "1111"

func main() {
	//var name string
	//var age init
	//var email string

	//var 1name string
	//var &test string
	//var *age int

	//if condition {
	//
	//}
	//
	//append

	//var name string
	//fmt.Printf("name: %\n", name)

	//var (
	//	name string
	//	age  int
	//	b    bool
	//)
	//
	//name = "tom"
	//age = 20
	//b = true

	//初始化 赋初值
	//var name string = "tom"
	//var site string = "duoko36.com"
	//var age int = 20

	//fmt.Printf("name: %v\n", name)
	//fmt.Printf("age: %v\n", age)
	//fmt.Printf("b: %v\n", b)

	//类型推断
	//var name = "tom"
	//var age = 20
	//var b = true
	//fmt.Printf("name: %v\n", name)
	//fmt.Printf("age: %v\n", age)
	//fmt.Printf("b: %v\n", b)

	//批量初始化
	//var name, age, b = "tom", 20, true
	//fmt.Printf("name: %v\n", name)
	//fmt.Printf("age: %v\n", age)
	//fmt.Printf("b: %v\n", b)

	//短变量声明
	//name := "tom"
	//age := 20
	//b := true
	//fmt.Printf("name: %v\n", name)
	//fmt.Printf("age: %v\n", age)
	//fmt.Printf("b: %v\n", b)

	name, age := getNameAndAge()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}

// 匿名变量
func getNameAndAge() (name string, age int) {
	return "tom", 20
}
