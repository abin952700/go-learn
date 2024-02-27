package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func initDb() {
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
	client = c
}

func main() {
	initDb()
}
