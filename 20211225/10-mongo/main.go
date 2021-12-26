package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.79.70:27017"))

	if err != nil {
		log.Errorf("conn mongo error:%s \v", err)
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// 创建数据库和集合
	collection := client.Database("testing").Collection("numbers")

	// 插入文档
	insertOne(&ctx, collection)

}

func insertOne(ctx *context.Context, collection *mongo.Collection) {
	res, err := collection.InsertOne(*ctx, bson.D{{"name", "pi"}, {"value", "3.14"}})
	if err != nil {
		log.Errorf("insert one error :%s\n", err)
		return
	}
	fmt.Printf("document id :%s\n", res.InsertedID)

}
