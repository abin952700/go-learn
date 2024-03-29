package main

import "fmt"

func h1() {
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
	fmt.Println("-------------")
	a1[0] = 1000
	a1[1] = 2000
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
	//数组长度越界
	//a1[3] = 1000
}

func h2() {
	// 数据的长度
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]int{1, 2, 3, 4}
	fmt.Printf("len(a1): %v\n", len(a2))
}

func h3() {
	//数组的遍历 1.根据和长度和下标
	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]：%v\n", i, a1[i])
	}
}

func h4() {
	//	for range
	var a1 = [3]int{1, 2, 3}
	for i, v := range a1 {
		fmt.Printf("a1[%v]：%v\n", i, v)
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	//h1()
	//h2()
	//h3()
	h4()
}
