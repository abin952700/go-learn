package main

import "fmt"

//切片是一个动态数组，可以使用append()函数添加元素，go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素，由于，切片是引用类型，通过赋值的方式，会修改原有内容,go提供了copy()函数来拷贝切片

// add
func l1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v/n", s1)
}

// del
func l2() {
	var s1 = []int{1, 2, 3, 4}
	//	i 2
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v/n", s1)
}

// update
func l3() {
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

// query
func l4() {
	var s1 = []int{1, 2, 3, 4}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

func l5() {
	var s1 = []int{1, 2, 3, 4}
	//var s2 = s1
	var s2 = make([]int, 4)
	copy(s2, s1) //s1是源
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func main() {
	//l1()
	//l2()
	//l3()
	//l4()
	l5()
}
