package main

import (
	"fmt"
	"strings"
)

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func main() {
	//var s string = "hello world"
	//var s1 = "Hello world"
	//s3 := "hello world"
	//
	//fmt.Printf("s: %v\n", s)
	//fmt.Printf("s1: %v\n", s1)
	//fmt.Printf("s3: %v\n", s3)
	//
	//s4 := `
	//line 1
	//line 2
	//line 3
	//`
	//fmt.Printf("s4: %v\n", s4)

	//字符串连接
	//s1 := "tom"
	//s2 := "20"
	//msg := s1 + s2
	//fmt.Printf("msg: %v\n", msg)

	//s1 := "tom"
	//s2 := "20"
	//msg := fmt.Sprintf("name=%s,age=%s", s1, s2)
	//fmt.Printf("msg: %v\n", msg)

	//name := "tom"
	//age := "20"
	//s := strings.Join([]string{name, age}, ",")
	//fmt.Printf("s: %v\n", s)

	//var buffer bytes.Buffer
	//buffer.WriteString("tom")
	//buffer.WriteString(",")
	//buffer.WriteString("20")
	//fmt.Printf("buffer.String(): %v\n", buffer.String())

	//转义字符
	//s := "hello \n world"
	//fmt.Sprintf("s: %v\n", s)
	//print(s + "\n")
	//print(s)

	//s := "Hello\tworld"
	//fmt.Printf("s: %v\n", s)

	//s := "I'm tom"
	//fmt.Printf("s: %v\n", s)

	//s := "c:\\program files\\a"
	//fmt.Printf("s：%v\n", s)

	//字符串的切片操作
	//s := "hello world"
	//a := 2
	//b := 5
	//fmt.Printf("s[a]: %c\n", s[a])     //取a位置的字符串
	//fmt.Printf("s[a,b]: %v\n", s[a:b]) //取a-b区间的字符串
	//fmt.Printf("s[a：]: %v\n", s[a:])   //取a开始到最后
	//fmt.Printf("s[:a]: %v\n", s[:b])   //取0开始到b位置的字符串

	//字符串函数
	s := "hello World"
	fmt.Printf("len(s): %v\n", len(s))

	//字符串分割
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))
	//字符串是否包含
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Contains(s, "hello"))
	//字符串转小写
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	//字符串转大写
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToUpper(s))
	//字符串是否以前缀结尾
	fmt.Printf("strings.HasPrefix(\"Hello\"): %v\n", strings.HasPrefix(s, "hello"))
	//字符串是否以后缀结尾
	fmt.Printf("strings.HasSuffix(\"world\"): %v\n", strings.HasSuffix(s, "world"))
	//字符串查找字符的位置
	fmt.Printf("strings.Index(\"ll\"): %v\n", strings.Index(s, "ll"))
	changeString()
}
