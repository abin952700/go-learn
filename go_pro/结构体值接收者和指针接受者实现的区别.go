package main

/*
结构体值接受者实现接口:
值接受者: 如果结构体中的方法是值接受者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口类型变量
*/

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

// 结构体的方法
//func (p Phone) start() { //值接收者
//	fmt.Plintln(p.Name, "启动")
//}
//
//func (p Phone) stop() {
//	fmt.Plintln(p.Name, "关机")
//}

func (p *Phone) start() { //指针接收者
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop() { //指针接收者
	fmt.Println(p.Name, "关机")
}

//定义一个Animal的接口，Animal中定义两个方法，分别是SerName和GetName,分别让Dog结构体和Cat实现这个方法
type Animaler interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

//引用类型
func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c Cat) GetName() string {
	return c.Name
}

func main() {
	// 结构体值接受者实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	//var p = Phone{
	//	Name: "小米手机",
	//}
	//
	//var p1 Usber = p //表示让Phone实现Usb的接口
	//p1.start()
	//
	//var p3 = &Phone{
	//	Name: "苹果手机",
	//}
	//var p4 Usber = p3 //表示让Phone实现Usb的接口
	//p4.start()

	//如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量，
	//结构体值类型没法赋值给接口变量。
	var phone1 = &Phone{
		Name: "小米",
	}
	var p1 Usber = phone1
	p1.stop()

	//Dog实现Animal的接口
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("阿奇")
	fmt.Println(d1.GetName())

	//Cat实现Animal的接口
	//&方法有指针类型接受者
	c := &Cat{
		Name: "小花",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())

}
