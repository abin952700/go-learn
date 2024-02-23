package main

import "fmt"

//golang匿名函数
//go语言函数不能嵌套，但是在函数内部可以定义匿名函数，实现一下简单功能调用
//func (参数列表)(返回值)

func t1() {
	//name := "tom"
	//age := 20
	//func() string {
	//
	//}
}

func main() {
	// 匿名函数 1
	//max := func(a int, b int) int {
	//	if a > b {
	//		return a
	//	} else {
	//		return b
	//	}
	//}
	//r := max(1, 2)
	//fmt.Printf("r: %v\n", r)

	//	匿名函數2 自己调用
	r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)
	fmt.Printf("r: %v\n", r)
}
