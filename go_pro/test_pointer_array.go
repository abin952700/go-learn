package main

import "fmt"

// var ptr [MAX]*int; 表示数组里面的元素类型是指针类型

func main() {
	a := [3]int{1, 2, 3}
	var pa [3]*int
	/*
	  pa = &a
	  pa ++
	  pa -- 指针不允许进行运算
	*/
	fmt.Printf("pa: %v\n", pa)
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}
	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v\n", *pa[i])
	}
}
