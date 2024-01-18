package main

import "fmt"

// break语句可以结束for、switch和select的代码块
// 使用break注意事项
// 1.单独在select中使用break和不使用没有啥区别
//2.单独在表达式switch语句，并且没有fallthough,使用break和不使用break没有啥区别
//3.单独在表达式switch语句，并且有fallthrough,使用break能够终止fallthrough后面的case语句的执行
//4.带标签的break,可以跳出多层select/switch作用域。让break更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带break的只能跳出当前语句块

func test3() {
MYLABEL:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break MYLABEL
		}
	}
	fmt.Println("END...")
}

func test2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		//break
		fallthrough
	case 3:
		fmt.Println("3")
	}
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

func main() {
	//test1()
	test2()
}
