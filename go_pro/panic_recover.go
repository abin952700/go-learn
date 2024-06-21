package main

import (
	"errors"
	"fmt"
)

/*
go 语言中目前是没有异常机制的，但是使用panic/recover 模式来处理错误
panic 可以在任何地方引发，但 recover 只有在defer 调用的函数中有限
*/

func fn() {
	defer func() {
		err := recover()
		if err != nil {
			//打印panic的异常
			fmt.Println("err:", err)
		}
	}()
	// 终止程序，并抛出错误
	panic("抛出错误")
}

func fn1(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error", err)
		}
	}()
	return a / b
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	//监听panic
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		//
		panic(err)
	}
}

func main() {
	//fmt.Println(fn1(10, 0)) //继续执行
	//fmt.Println("结束")

	myFn()
	fmt.Println("继续执行")
}
