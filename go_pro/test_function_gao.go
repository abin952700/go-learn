package main

import "fmt"

// go语言的函数，可以作为函数的参数，传递给另外一个函数，可以作为，另外一个函数的返回值返回

func sayHello(name string) {
	fmt.Printf("Hello,%s", name)
}

// 函数传入
func s1(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// 返回值是一个函数
func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	s1("tom", sayHello)
	ff := cal("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)
	ff = cal("-")
	r = ff(2, 1)
	fmt.Printf("r: %v\n", r)
}
