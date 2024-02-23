package main

import "fmt"

// 数组是固定长度，可以容纳相同数据类型的元素的集合。当长度固定时，使用还是带来一些限制。比如：我们申请的长度太大浪费内存，太小又不够用。
//鉴于上述原因，我们有了go语言的切片，可以把切片理解为，可变数组的长度，其实它底层就是使用数组实现的，增加了自动扩容功能。切片（slice）是一个拥有相同类型元素的可变长度的序列。

func i1() {
	//var a1 = [2]int{1, 2, 3}
	//append(a1,100)
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func i2() {
	// 通过make声明切片 初始化
	var s2 = make([]int, 2)
	fmt.Printf("s2：%v\n", s2)
}

func i3() {
	var s1 = []int{1, 2, 3}
	fmt.Printf("len(s1)：%v\n", len(s1))
	//容量
	fmt.Printf("cap(s1)：%v\n", cap(s1))
	fmt.Printf("s1[0]: %v\n", s1[0])
}

func main() {
	//i1()
	//i2()
	i3()
}
