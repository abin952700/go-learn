package main

import "fmt"

func f1_() {
	score := 90
	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func f2_() {
	if score := 90; score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func f3_() {
	var c string
	fmt.Println("请输入一个字符：")
	fmt.Scan(&c)
	if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuseday")
		} else {
			fmt.Println("Thursday")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else if c == "F" {
		fmt.Println("Firday")
	} else if c == "S" {
		fmt.Println("请输入第二个字符： ")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Staturday")
		} else {
			fmt.Println("Sunday")
		}
	}

}

func main() {
	//f1_()
	//f2_()
	f3_()
}
