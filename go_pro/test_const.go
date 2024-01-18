package main

import "fmt"

func main() {
	//const PI float32 = 3.14
	////PI = 3.15
	//const PI2 = 3.1415
	//fmt.Printf("PI: %\n", PI)
	//
	//const (
	//	a = 100
	//	b = 100
	//)
	//fmt.Printf("a: %v\n", a)
	//fmt.Printf("b: %v\n", b)
	//
	//const w, h = 200, 300
	//fmt.Printf("w:%v\n", w)
	//fmt.Printf("h:%v\n", h)

	//iota
	//const (
	//	a1 = iota
	//	//a2 = iota
	//	_
	//	a3 = iota
	//)
	//fmt.Printf("a1: %v\n", a1)
	//fmt.Printf("a2: %v\n", a2)
	//fmt.Printf("a3: %v\n", a3)

	const (
		a1 = iota
		a2 = 100 //跳过
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a3: %v\n", a3)
}
