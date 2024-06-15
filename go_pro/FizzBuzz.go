package main

import "fmt"

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func reverseString(s string) string {
	//	将字符串转换为字节数组
	str := []rune(s)

	// 定义两指针，一个指向字符串的开头，一个指向结尾
	left, right := 0, len(str)-1

	//	交换字符直到指针相遇
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	//	将字节数组转换回字符串返回
	return string(str)
}

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciRecursive(i))
	}

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciIterative(i))
	}

	s := "Hello, world!"
	fmt.Println(reverseString(s))
}
