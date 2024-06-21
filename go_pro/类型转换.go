package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  数值类型之间的转换

	//1.整型和整型之间的转换
	//var a int8 = 20
	//var b int16 = 40
	//fmt.Println(int16(a)+b)

	//2.浮点型和浮点型之间的转换
	//var a float32 = 20
	//var b float64 = 40
	//fmt.Println(float64(a) + b)

	//3.整形和浮点型之间的转换
	//var a float32 = 20.32
	//var b int = 40
	//fmt.Println(a + float32(b))

	//注意：转换的时候建议从低位转换成高位，高位转换成低位的时候如果不成功就会溢出，和我们想的结果不一样
	//var a1 int8 = 20
	//var b1 int16 = 140

	//其他类型转换成string类型

	//注意：sprintf使用中需要注意转换的格式int为%d float为%f bool为%t byte为%c

	//var i int = 20
	//var f float64 = 12.456
	//var t bool = true
	//var b byte = 'a'
	////1.sprintf把其他类型转换成string类型
	//str1 := fmt.Sprintf("%d", i)
	//fmt.Printf("值:%v--类型:%T", str1, str1)
	////保留两位小数
	//str2 := fmt.Sprintf("%.2f", f)
	//fmt.Printf("值:%v--类型:%T", str2, str2)
	//
	//str3 := fmt.Sprintf("%t", t)
	//fmt.Printf("值:%v--类型:%T", str3, str3)
	//
	//str4 := fmt.Sprintf("%t", b)
	//fmt.Printf("值:%v--类型:%T", str4, str4)

	//2.使用strconv包里面的几种转换方法进行转换
	/*
			  FormatInt
			  参数1： int64的数值
			  参数2： 传值int类型的进制

			  FormatFloat
			  参数1：要转换的值
			  参数2：格式化类型'f'（-ddd.dddd）
			  参数3：保留的小数点 -1(不对小数点格式化)
		      参数4：格式化的类型 传入 64 32


	*/
	//var i int = 20
	//str1 := strconv.FormatInt(int64(i), 10)
	//fmt.Println(str1)

	//var f float32 = 20.23
	//str2 := strconv.FormatFloat(float64(f), 'f', 2, 32)
	//fmt.Printf("值:%v--类型:%T", str2, str2)
	//
	//str3 := strconv.FormatBool(true)
	//fmt.Printf("值:%v--类型:%T", str3, str3)
	//
	//a := 'a'
	//str4 := strconv.FormatUint(uint64(a), 10)
	//fmt.Printf("值:%v--类型:%T", str4, str4)

	//string类型转换成整形
	str := "123456"
	fmt.Printf("%v----%T", str, str)

	/*
	  parseInt
	  参数1：string
	  参数2：进制
	  参数3：位数 32 64 16
	*/
	num, _ := strconv.ParseInt(str, 10, 64)
	num1, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("%v----%T", num, num)
	fmt.Printf("%v----%T", num1, num1)

	//	在go语言中数值类型设法直接转换成bool类型 bool类型也设法直接转换成数值类型

	// 不建议string 转 bool 非空也是bool形
	b, _ := strconv.ParseBool("xxxxx")
	fmt.Printf("%v----%T", b, b)
}
