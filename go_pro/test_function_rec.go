package main

import "fmt"

// 使用递归函数的3要点：
/*
1.递归就是自己调用自己
2.必须先定义函数的退出条件。没有退出条件，递归将成为死递归
3.go语言递归函数可能会产生一大堆的gorouline,也很可能会出现栈空间内存溢出的问题
*/

//func u1() {
//	fmt.Printf("...")
//	u1()
//}

func u2() {
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

func u3(a int) int {
	if a == 1 {
		// 1.退出条件
		return 1
	} else {
		// 2.自己调用自己
		return a * u3(a-1)
	}
}

func main() {
	u2()
	i := u3(6)
	fmt.Printf("i: %v\n", i)
}
