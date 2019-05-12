package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

// startTime小于某时间
// {"$lt": timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

// {"timePoint.startTime": {"$lt": timestamp} }
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

var (
	client     *mongo.Client
	err        error
	database   *mongo.Database
	collection *mongo.Collection

	resultOne  *mongo.InsertOneResult
	resultMany *mongo.InsertManyResult

	id objectid.ObjectID

	logArr []interface{}
)

func getConnect() *mongo.Client {

	// 建立连接
	//var mongo_db_url = "mongodb://192.168.31.152:27017"
	var mongo_db_url = "mongodb://111.231.77.158:27017"
	if client, err = mongo.Connect(context.TODO(), mongo_db_url, clientopt.ConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func delete() {

	var (
		delCond   *DeleteCond
		delResult *mongo.DeleteResult
	)

	client = getConnect()

	// 选择数据库
	database = client.Database("cron")
	// 选择表
	collection = database.Collection("log")

	// 要删除开始时间早于当前时间的所有日志($lt是less than)
	//  delete({"timePoint.startTime": {"$lt": 当前时间}})
	delCond = &DeleteCond{beforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	// 执行删除
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数:", delResult.DeletedCount)

}

func main() {
	delete()
}
