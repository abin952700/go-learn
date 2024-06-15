package main

import "fmt"

//type student struct {
//	name string
//	age  int
//}

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

func main() {
	//m := make(map[string]*student)
	//stus := []student{
	//	{name: "pprof.cn", age: 18},
	//	{name: "测试", age: 23},
	//	{name: "博客", age: 28},
	//}
	//
	//for _, stu := range stus {
	//	m[stu.name] = &stu
	//}
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.name)
	//}

	//p9 := newPerson("pprof.cn", "测试", 90)
	//fmt.Printf("%#v\n", p9)

	//p1 := NewPerson("测试", 25)
	//fmt.Println(p1.age) // 25
	//p1.SetAge(30)
	//fmt.Println(p1.age) // 30

	//d1 := &Dog{
	//	Feet: 4,
	//	Animal: &Animal{ //注意嵌套的是结构体指针
	//		name: "乐乐",
	//	},
	//}
	//d1.wang() //乐乐会汪汪汪~
	//d1.move() //乐乐会动！

	//var ce []student //定义一个切片类型的结构体
	//ce = []student{
	//	student{1, "xiaoming", 22},
	//	student{2, "xiaozhang", 33},
	//}
	//fmt.Println(ce)
	//demo(ce)
	//fmt.Println(ce)

	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}
