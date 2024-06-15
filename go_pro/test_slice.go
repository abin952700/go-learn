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
	//i3()
	//var s3 []int = make([]int, 0)
	//fmt.Println(s3)
	//// 4.初始化赋值
	//var s4 []int = make([]int, 0, 0)
	//fmt.Println(s4)
	//s5 := []int{1, 2, 3}
	//fmt.Println(s5)
	//// 5.从数组切片
	//arr := [5]int{1, 2, 3, 4, 5}
	//var s6 []int
	//// 前包后不包
	//s6 = arr[1:4]
	//fmt.Println(s6)
	//
	//arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//slice8 := arr2[:]
	//fmt.Println(slice8)
	//
	//s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	//fmt.Println(s1, len(s1), cap(s1))
	//
	//s := []int{0, 1, 2, 3}
	//p := &s[2] // *int, 获取底层数组元素指针。
	//*p += 100
	//
	//fmt.Println(s)

	//d := [5]struct {
	//	x int
	//}{}
	//
	//s := d[:]
	//
	//d[1].x = 10
	//s[2].x = 20
	//
	//fmt.Println(d)
	//fmt.Printf("%p, %p\n", &d, &d[0])

	//data := [...]int{0, 1, 2, 3, 4, 10: 0}
	//s := data[:2:3]
	//
	//s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。
	//
	//fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	//fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。

	// 创建一个切片，初始容量为 2
	//slice := make([]int, 0, 2)
	//fmt.Printf("初始容量：%d\n", cap(slice))
	//
	//// 添加元素，直到切片容量不足
	//for i := 0; i < 10; i++ {
	//	slice = append(slice, i)
	//	fmt.Printf("长度：%d, 容量：%d\n", len(slice), cap(slice))
	//}

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)
}
