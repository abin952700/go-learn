package main

import "fmt"

//golang有一个特殊的函数init函数，先于main函数执行，实现包级别的一些初始化操作

/*
init函数的主要特点
1.init函数先于main函数自动执行，不能被其他函数调用
2.init函数没有输入参数，返回值
3.每个包可以有多个init函数
4.包的每个源文件也可以有多个init函数，这点比较特殊
5.同一个包的init执行顺序，golang没有明确定义，编辑时要注意程序不要依赖这个执行顺序
6.不同包的init函数按照导入的依赖关系决定执行顺序
*/

//golang初始化顺序 变量初始化->init()->main

var i int = initVar()

var j int

func init() {
	fmt.Println("init...")
	j = 100
}

func init() {
	fmt.Println("init2...")
}

func initVar() int {
	fmt.Println("initVar...")
	return 100
}

func main() {
	fmt.Println("main...")
}
