package main

import "fmt"

func b1() {
	//for i := 1; i <= 10; i++ {
	//	fmt.Printf("i: %v\n", i)
	//}
}

func b2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func b3() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

func b4() {
	for {
		fmt.Printf("run...")
	}
}

func main() {
	//b1()
	//b2()
	//b3()
	b4()
}
