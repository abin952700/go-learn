package main

import "fmt"

//切片初始化
//切片的初始化方法很多,可以直接初始化,也可以数组初始化等

// 切片
func j1() {
	//初始化方式一
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:3] //[ )
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:] //从3取到最后
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
	s5 := s1[:]
	fmt.Printf("s4: %v\n", s5)
}

// 数组
func j2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
	a3 := a1[3:]
	fmt.Printf("a3: %v\n", a3)
	a4 := a1[:]
	fmt.Printf("a4: %v\n", a4)
}

func main() {
	//j1()
	//j2()
}
