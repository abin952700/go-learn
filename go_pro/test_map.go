package main

import (
	"fmt"
)

//map是一种key:value键值对的数据结构容器。map内部实现是哈希表(hash).map最重要的一点是通过key来快速检索数据。key类似于索引，指向数据的值.map是引用类型的。

func m1() {
	var m1 map[string]string     // 1 声明变量，默认map是nil
	m1 = make(map[string]string) //2 使用make函数
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

// 初始化
func m2() {
	//map是无序的
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	fmt.Printf("m1: %v\n", m1)
	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["email"] = "tom@gmail.com"
	fmt.Printf("m2: %v\n", m2)
}

func m3() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}

// 判断某个键是否存在
func m4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v： %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("-------------")
	v, ok = m1[k2]
	fmt.Printf("v： %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}

func main() {
	//m1()
	//m2()
	//m3()
	//m4()
	//scoreMap := make(map[string]int)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//scoreMap["王五"] = 60
	//for k := range scoreMap {
	//	fmt.Println(k)
	//}
	//scoreMap := make(map[string]int)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//scoreMap["王五"] = 60
	//delete(scoreMap, "小明") //将小明:100从map中删除
	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}

	//rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	//
	//var scoreMap = make(map[string]int, 200)
	//
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	//	value := rand.Intn(100)          //生成0~99的随机整数
	//	scoreMap[key] = value
	//}
	////取出map中的所有key存入切片keys
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	////对切片进行排序
	//sort.Strings(keys)
	////按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

	//var mapSlice = make([]map[string]string, 3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "王五"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "红旗大街"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		fmt.Println(111111111)
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
