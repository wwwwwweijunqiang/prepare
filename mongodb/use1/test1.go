package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

func main() {

	var (
		client   *mongo.Client
		err      error
		database *mongo.Database

		collection *mongo.Collection
	)

	// 建立连接
	var mongo_db_url = "mongodb://192.168.31.6:27017"
	if client, err = mongo.Connect(context.TODO(), mongo_db_url, clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}
	// 选择数据库 my_db
	database = client.Database("my_db")
	// 选择表 my_collection

	collection = database.Collection("my_collection")

	collection = collection
	fmt.Println("111")
}
