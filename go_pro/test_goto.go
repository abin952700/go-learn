package main

import "fmt"

//goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环，避免重复退出上有一定帮助。
//Go语言中使用goto语句能简化一些代码的实现过程。例如双层嵌套的for循环要退出时：

func e2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END
			}
			fmt.Printf("%v,%v", i, j)
		}
	}
END:
	fmt.Printf("END...")
}

func e1() {
	i := 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}
END:
	fmt.Println("END.....")
}

func main() {
	e1()
}
