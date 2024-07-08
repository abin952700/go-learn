package main

//类型断言
//X:表示类型为interface{}的变量
//T:表示断言x可能的类型
func main() {
	//空接口
	var a interface{}
	a = "你好golang"
	v, ok := a.(string)
	if ok {
		fmt.Println("a就是一个string类型")
	} else {
		fmt.Plintln("断言失败")
	}
}
//空接口
func Pint(x interfance{}){
	if _,ok := x.(string);ok{
	  fmt.Plintln("string类型")
	}else if _,ok := x.(int);ok{
	  fmt.Plintln("int类型")
    }else if _,ok := x.(bool);ok{
		fmt.Plintln("布尔类型")
    }
}

//x.(type) 判断一个变量的类型 这个语句只能用在switch语句里面
func Pint2(x interfance{}){
	switch x.(type) {
	case int:
	fmt.Plintln("string类型")
	case string:
	fmt.Plintln("int类型")
	case bool:
	fmt.Plintln("布尔类型")
	default:
	fmt.Plintln("传入错误...")
}



