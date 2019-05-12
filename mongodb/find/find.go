package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
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

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"`
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
func findOne() {

	var (
		cond *FindByJobName

		cursor mongo.Cursor

		record *LogRecord
	)
	client = getConnect()
	fmt.Println(11111)
	// 选择数据库
	database = client.Database("cron")
	// 选择表
	collection = database.Collection("log")
	// 过滤条件
	cond = &FindByJobName{JobName: "job110"}
	// 查询(过滤+分页)
	if cursor, err = collection.Find(context.TODO(), cond, findopt.Skip(0), findopt.Limit(2)); err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())
	fmt.Println(22222)
	// 遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义日志对象
		record = &LogRecord{}
		// bson数据反序列化为LogRecord类型
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*record)
	}

}
func main() {
	findOne()
}
