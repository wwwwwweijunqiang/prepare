package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error

	putOP clientv3.Op
	kv    clientv3.KV
)

func getConnect() *clientv3.Client {

	// 配置参数
	config = clientv3.Config{
		Endpoints:   []string{"192.168.31.6:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func operations() {

	client = getConnect()
	kv = clientv3.NewKV(client)
	// 创建operation
	putOP = clientv3.OpPut("/cron/jobs/job8", "this is op job8")
	// 执行op
	if opResp, err := kv.Do(context.TODO(), putOP); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("写入Revision: ", opResp.Put().Header.Revision)
	}

	// 创建operation
	getOP := clientv3.OpGet("/cron/jobs/job8")
	// 执行op
	if opResp, err := kv.Do(context.TODO(), getOP); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("数据Revision: ", opResp.Get().Kvs[0].ModRevision)
		fmt.Println("数据value: ", string(opResp.Get().Kvs[0].Value))
	}

}

func main() {
	operations()
}
