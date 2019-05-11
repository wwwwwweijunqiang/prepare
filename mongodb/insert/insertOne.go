package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

var (
	client     *mongo.Client
	err        error
	database   *mongo.Database
	collection *mongo.Collection
	record     *LogRecord

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

func insertOne() {
	client = getConnect()
	// 选择数据库
	database = client.Database("cron")
	// 选择表
	collection = database.Collection("log")
	fmt.Println("选择表")
	record = &LogRecord{
		JobName:   "job110",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	if resultOne, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}
	//  _id:，默认生成一个全局唯一id，objectID:12字节的二进制
	id = resultOne.InsertedID.(objectid.ObjectID) // 类型转换（反射）
	fmt.Println("自增id： ", id.Hex())
}

func insertMany() {

	client = getConnect()
	// 选择数据库
	database = client.Database("cron")
	// 选择表
	collection = database.Collection("log")
	fmt.Println("选择表")
	record = &LogRecord{
		JobName:   "job110",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	logArr = []interface{}{record, record, record}

	resultMany, err = collection.InsertMany(context.TODO(), logArr)
	if err != nil {
		fmt.Println("insert many err")
		fmt.Println(err)
		return
	}
	for _, insertId := range resultMany.InsertedIDs {
		// 把insertId反射成objectID
		id := insertId.(objectid.ObjectID)
		fmt.Println("自增id: ", id)
	}

}

func main() {

	insertMany()
}
