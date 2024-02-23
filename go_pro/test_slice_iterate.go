package main

import "fmt"

func k1() {
	var s1 = []int{1, 2, 3, 4, 5}
	l := len(s1)
	for i := 0; i < l; i++ {
		fmt.Printf("s1[%v]：%v\n", i, s1[i])
	}
}

func k2() {
	var s1 = []int{1, 2, 3, 4, 5}
	for i, v := range s1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	test1()
	test2()
}
