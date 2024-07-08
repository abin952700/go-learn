package main

type Animaler1 interface {
	SetName(string)
}

type Animaler2 interface {
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

func main() {

	//Dog实现Animal的接口
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animaler1 = d //表示让dog实现Animaler1这个接口
	var d2 Animaler2 = d //表示让dog实现Animaler2这个接口

	d1.SetName("小花")
	fmt.Plintln(d2.GetName())

}
