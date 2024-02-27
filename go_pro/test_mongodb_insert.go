package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//创建一个结构体
/*
type Student struct {
  Name string
  Age int
}
*/
// 添加单个文档
// 使用collection.InsertOne()方法插入一条文档记录：

type Student struct {
	Name string
	Age  int
}

var client1 *mongo.Client

func initDb1() {
	// 设置客户端连接配置
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	c, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("c: %v\n", c)
	}
	// 检测连接
	err2 := c.Ping(context.TODO(), nil)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}
	client1 = c
}

func insertMany() {
	// 初始化数据库
	initDb1()
	c := client1.Database("go_db").Collection("Student")

	s1 := Student{
		Name: "kiite",
		Age:  20,
	}

	s2 := Student{
		Name: "rose",
		Age:  20,
	}

	stus := []interface{}{s1, s2}
	imr, err := c.InsertMany(context.TODO(), stus)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("imr.InsertedID: %v\n", imr.InsertedIDs)
}
func insert() {
	initDb1()
	s1 := Student{
		Name: "tom",
		Age:  20,
	}
	c := client1.Database("go_db").Collection("Student")
	ior, err := c.InsertOne(context.TODO(), s1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
	}
}

func main() {
	//insert()
	insertMany()
}
