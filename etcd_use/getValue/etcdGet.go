package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	getResp *clientv3.GetResponse
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

// 得到具体路径下面的值
func getValue() {

	client = getConnect()
	// 读写etcd的键值对
	kv = clientv3.NewKV(client)
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(getResp.Kvs)
	}

}

// 得到/cron/jobs路径下面的值
func getAllValue() {
	client = getConnect()
	// 读写etcd的键值对
	kv = clientv3.NewKV(client)

	kv.Put(context.TODO(), "/cron/jobs/job2", "job2")
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		// 获取成功，查看所有数据

		fmt.Println(getResp.Kvs)
	}

}

func main() {
	//getAllValue()
	getValue()
}
