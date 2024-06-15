package main

import "fmt"

//go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无需拷贝数据。
// 类型指针不能使用偏移和运算
//go语言中的指针操作非常简单，只需要记住两个符号&（取地址）和 *（根据地址取值）

/*
指针地址和指针类型
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。go语言中使用&字符放在变量前面对变量进行取地址的操作。Go语言中值类型（int、float、bool、string、array、struct）都有对应的指针类型。如*int、*int64、*string等


指针语法
一个指针变量指向了一个值的内存地址。（也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中。）
类似于变量和常量，在使用指针前你需要声明指引。指针声明格式如下：
var var_name *var-type
var-type :为指针类型
var_name :指针变量名

* ： 用于指定变量是作为一个指针
*/

func main() {
	var ip *int                //声明一个指针类型
	fmt.Printf("ip: %v\n", ip) //nil
	fmt.Printf("ip: %T\n", ip) //int类型的指针类型
	var i int = 100
	ip = &i                     //给指针赋值
	fmt.Printf("ip: %v\n", ip)  //取地址
	fmt.Printf("ip: %v\n", *ip) //取值

	var sp *string
	var s string = "hello"
	sp = &s
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %v\n", *sp)

	var bp *bool
	var b bool = true
	fmt.Printf("bp: %T\n", bp)
	bp = &b //取地址对应指针类型
	fmt.Printf("bp: %v\n", bp)
	fmt.Printf("bp: %v\n", *bp)

	//var p *string
	//fmt.Println(p)
	//fmt.Printf("p的值是%s/n", p)
	//if p != nil {
	//	fmt.Println("非空")
	//} else {
	//	fmt.Println("空值")
	//}

	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)

}
