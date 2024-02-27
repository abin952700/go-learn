package main

import "fmt"

//类型定义语法

// type NewType Type
func main() {
	//type MyInt int
	//var i MyInt
	//i = 100
	//fmt.Printf("i: %T,%v\n", i, i)

	//类型别名的语法
	//type NewType = Type
	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("i: %T, %v\n", i, i)
}

//go语言类型定义和类型别名的区别
/*
1.类型定义相当于定义了一个全新的类型,与之前那的类型不同；但是类型别名并没有定义一个新的类型，而是使用一个别名来替换之前的类型
2.类型别名只会在代码中存在，在编译完成之后并不会存在该别名
3.因为类型别名和原来的类型是一致的，所以原来所拥有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前额度任何方法
*/
