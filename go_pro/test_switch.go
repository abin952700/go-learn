package main

import "fmt"

func a1() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("其他")
	}
}

func a2() {
	day := 2
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
	}
}

func a3() {
	score := 60
	switch {
	case score >= 60:
		fmt.Println("及格")
	case score >= 80 && score < 90:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}
}

func f4() {
	num := 100
	switch num {
	case 100:
		fmt.Println("100")
		fallthrough //穿透
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	}
}

func main() {
	//a1()
	//a2()
	a3()
}

//go语言中switch语句的注意事项
//1.支持多条件匹配
//2.不同的case之间不使用break分割，默认之后执行一个case
//3.如果想要执行多个case,需要使用fallthrough关键字，也可以break终止
