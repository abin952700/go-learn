package main

import "fmt"

func f1__() {
	a, b, c := 100, 2, 3
	if a > b {
		if a > c {
			fmt.Println("a")
		}
	} else {
		// b>a
		if b > c {
			fmt.Println("b")
		} else {
			fmt.Println("c")
		}
	}
}

func f2__() {
	gender := "男"
	age := 20
	if gender == "男" {
		if age >= 18 {
			fmt.Println("成年男性")
		} else {
			fmt.Println("未成年男性")
		}
	} else {
		if age >= 18 {
			fmt.Println("成年女性")
		} else {
			fmt.Println("未成年女性")
		}
	}
}
func main() {
	f2__()
}
