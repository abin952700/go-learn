package main

import "fmt"

/*
go语言函数可以有0或多个参数，参数需要指定的数据类型

声明函数时的参数列表叫形参，调用时传递的参数叫实参

go语言是通过传值的方式传参的，意味着传递给函数的是拷贝后的副本，所以函数内部访问、修改的也是这个副本。

go语言可以使用变长参数，有时候并不能确定参数的个数，可以使用变长参数，可以在函数定义语句的参数部分使用ARGS...TYPE的方式。这时会将...代表的参数全部保存到一个名为ARGS的slice中，注意这些参数的数据类型都是TYPE
*/

func q1(a int, b int) int {
	return a + b
}

func q2(a int) {
	a = 100
}

func q3(s []int) {
	s[0] = 1000
}

// 可变参数 1
func q4(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

// 可变参数 2
func q5(name string, ok bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("ok: %v\n", ok)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	r := q1(1, 2)
	fmt.Printf("r: %v\n", r)
	x := 200
	//传递是一个副本
	q2(x)
	fmt.Printf("x: %v\n", x)
	//map、slice、interface\channel这些数据类型本身就是"指针"类型，所以就算是拷贝传值也是拷贝的指针，拷贝后的参数任然指向底层数据结构，所以修改他们可能会影响外部数据结构的值。
	s := []int{1, 2, 3}
	q3(s)
	fmt.Printf("s: %v\n", s)
	q4(1, 2, 3, 4)
	q5("tom", true, 1, 2, 3)
}
