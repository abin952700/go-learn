package main

import "fmt"

// continue只能用在循环中，在go中只能用在for循环中，它可以终止本次循环，进行下一次循环。
//在continue语句后添加标签时，表示开始标签对应的循环

func d2() {
	for i := 0; i < 10; i++ {
	MYLABEL:
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				continue MYLABEL
			}
			fmt.Printf("%v, %v\n", i, j)
		}
	}
}

func d1() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("i: %v\n", i)
		} else {
			continue
		}

	}
}

func main() {

}
