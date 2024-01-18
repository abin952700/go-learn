package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	//格式化输出

	// %v var 相应值的默认格式
	//fmt.Printf("\"hello\": %v\n", "hello")

	// 定义结构体变量
	//site := WebSite{Name: "duoke360"}
	//// %#v 输出整个结构
	////fmt.Printf("site: %#v\n", site)
	//
	//// %T 输出类型
	//fmt.Printf("site: %T\n", site)
	//a := 100
	//fmt.Printf("a: %T\n", a)
	//fmt.Printf("%%")
	//
	//b := true
	//// %t布尔类型的输出
	//fmt.Printf("b: %t\n", b)

	i := 96
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)    //二进制8
	fmt.Printf("i: %c\n", i)    //对应unicod码所表示的字符
	fmt.Printf("i: %x\n", 100)  //100转换成16进制的输出
	fmt.Printf("i: %x\n", 1234) //1234转换成16进制的输出

	x := 100
	p := &x
	fmt.Printf("i: %p\n", p) //指针类型
}
