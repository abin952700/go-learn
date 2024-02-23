package main

import "fmt"

/*
函数可以有0个或多个返回值，返回值需要指定数据类型，返回值通过return关键字来指定。
return可以有参数，也可以没有参数，这些返回值可以有名称，也可以没有名称。go中的函数可以有多个返回值。

1.return关键字中指定了参数时，返回值可以不用名称。如果return省略参数，则返回值部分必须带名称。
2.当返回值有名称时，必须使用括号包围，逗号分隔，即使只是一个返回值
3.但即使返回值命名了，return中也可以强制制定其他返回值的名称，也就是说return的优先级更高
4.命名的返回值是预先声明好的，在函数内部可以直接使用，无需再次声明。命名返回值的名称不能和函数参数名称相同，否则报错提示变量重复定义
5.return中可以有表达式，但不能出现赋值表达式，这和其他语言可能有所不同，例如return a+b是正确的，但return c-a+b是错误的。
*/

// 没有返回值
func p1() {
	fmt.Printf("我既没有参数，也没有返回值...")
}

// 有一个返回值 1
func sum1(a int, b int) (ret int) {
	ret = a + b
	return ret
}

// 有一个返回值 2
func sum2(a int, b int) int {
	//ret := a + b
	//return ret
	return a + b
}

// 多个返回值，且在return中指定返回的内容
func p2() (name string, age int) {
	name = "老郭"
	age = 30
	return name, age
}

// 多个返回值，返回值名称没有被使用
func p3() (name string, age int) {
	name = "tom"
	age = 20
	//return name, age  等价于
	return
}

// return覆盖命名返回值，返回值名称没有被使用
func p4() (string, int) {
	n := "老--"
	a := 30
	return n, a
}

/*
go中经常会使用其中一个返回值作为函数是否执行成功，是否有错误信息的判断条件。列如return value,exists、 return value,ok, return value,err等

当函数的返回值过多时，列如有4个以上的返回值，应该将这些返回值收集到容器中，然后以返回容器的方式去返回。例如，同类型的返回值可以放进slice中，不同类型的返回值可以放进map中。

但函数有多个返回值时，如果其中某个或某几个返回值不想使用，可以通过下划线_来丢弃这些返回值。
*/

func p5() (int, int) {
	return 1, 2
}

func main() {
	name, age := p2()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	_, x := p5()
	fmt.Printf("x: %v\n", x)
}
