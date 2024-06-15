package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组是相同数据类型的一组数据的集合，数组一旦定义长度不能修改，数据可以通过下标（或者叫索引）来访问元素。
//数组定义 var variable_name [SIZE] variable_type

func g1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [2]bool
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	// 数组的初始化
	// 初始化，就是给数组的元素赋值，没有初始化的数组，就认元素值都是零值，布尔类型是false，字符串是空字符串
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}

func g2() {
	//数组的初始化
	//初始化列表
	var a1 = [3]int{1, 2}
	fmt.Printf("a1：%v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2：%v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3：%v\n", a3)
}

func g3() {
	//数组的初始化，默认长度省略长度 ...
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))
	var a2 = [...]string{"duoko360.com", "golang"}
	fmt.Printf("a2: %v\n", a2)
}

func g4() {
	var a1 = [...]int{0: 1, 3: 100, 5: 10}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]bool{2: true, 5: false}
	fmt.Printf("a2: %v\n", a2)
}

func g5() {
	var a1 = [...]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	a1[0] = 100
	fmt.Printf("a1 %v\n", a1)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	//g1()
	//g2()
	//g3()
	//g4()
	//g5()

	//var arr1 [5]int
	//printArr(&arr1)
	//fmt.Println(arr1)
	//arr2 := [...]int{2, 4, 6, 8, 10}
	//printArr(&arr2)
	//fmt.Println(arr2)

	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
}
