package main

import "fmt"

//go语言中函数的特性
// 函数是go语言中的一级公民
/*
1.go语言中有3种函数：普通函数、匿名函数、方法（定义在struct上的函数） receiver
2.go语言中不允许函数重载，也就是说不允许函数同名。
3.go语言中的函数不能嵌套函数，但可以嵌套匿名函数。
4.函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数。
5.函数可以作为参数传递给另一个函数。
6.函数的返回值可以是一个函数。
7.函数调用的时候，如果有参数表传递给函数，则先拷贝函数的副本，再将副本传递给函数。
8.函数参数可以没有名称。
*/

// 函数定义
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 函数定义
func comp(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

// 函数调用
func main() {
	r := sum(1, 2)
	fmt.Printf("r: %v\n", r)
	r = comp(1, 2)
	fmt.Printf("r: %v\n", r)
}
