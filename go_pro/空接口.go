package main

type A interface{} //空接口 表示没有任何约束 任意的类型都可以实现空接口

func main() {
	var a A
	var str = "hello"
	a = str //让字符串实现a这个接口
	fmt.Printf("值%v 类型:%T\n", a, a)

	var num = 20
	a = num //表示让int类型实现A这个接口
	fmt.Printf("值%v 类型:%T\n", a, a)

	var flag = true
	a = flag //表示让bool类型实现A这个接口
	fmt.Printf("值%v 类型:%T", a, a)
	show(20)
	show("你好")
	show([]int{1, 2, 3, 4, 5})
	//map类型
	var m1 = make(map[string]string)
	var m2 = make(map[string]interface{})

	m1["username"] = "张三"
	m1["age"] = "20"
	fmt.Plintln(m1)
	fmt.Plintln(m2)
	//切片类型
	var s1 = []int{1, 2, 3, 4}
	var s2 = []interface{}{1, "2", true}
}

//golang中空接口也可以直接当做类型来使用，可以表示任意类型
func test() {
	var a interface{}
	a = 20
	a = true
}
func show(a interface{}) {
	fmt.Printf("值%v 类型:%T", a, a)
}
