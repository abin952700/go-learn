package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client2 *mongo.Client

func initDb2() {
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
	client2 = c
}

func find() {
	ctx := context.TODO()
	// 关闭数据库
	defer func(client2 *mongo.Client, ctx context.Context) {
		err := client2.Disconnect(ctx)
		if err != nil {
			fmt.Printf("%v 连接错误", err)
		}
	}(client2, ctx)
	c := client2.Database("go_db").Collection("Student")
	c2, err := c.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer func(c2 *mongo.Cursor, ctx context.Context) {
		err := c2.Close(ctx)
		if err != nil {
			fmt.Printf("%v 关闭错误", err)
		}
	}(c2, ctx)

	for c2.Next(ctx) {
		var result bson.D
		c2.Decode(&result)
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result: %v\n", result.Map())
	}
}

func main() {
	initDb2()
	find()
}
