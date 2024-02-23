package main

import "fmt"

//闭包可以理解成定义一个函数内部的函数。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说。闭包 = 函数 + 引用环境。

func add1() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal1(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}

	sub := func(a int) int {
		base -= a
		return base
	}

	return add, sub
}

func main() {
	f := add1()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("---------")
	f1 := add1()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	add, sub := cal1(100)
	r = add(100)
	fmt.Printf("r: %v\n", r)
	r = sub(50)
	fmt.Printf("r: %v\n", r)
}
