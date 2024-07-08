package main

import "fmt"

type Adderss struct {
	Name  string
	Phone int
}

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])
	//fmt.Println(userinfo["hobby"][1]) //interface{} does not support indexing

	var address = Adderss{
		Name:  "李四",
		Phone: 1525163123,
	}
	userinfo["address"] = address
	fmt.Println(userinfo["address"])
	//var name = userinfo["address"].Name //type interface{} is interface with so
}
