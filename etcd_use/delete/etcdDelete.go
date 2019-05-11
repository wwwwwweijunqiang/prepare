package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
	kv     clientv3.KV

	delResp *clientv3.DeleteResponse
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

func deleteValue() {

	client = getConnect()
	log.Warn("1")

	// 读写etcd的键值对
	kv = clientv3.NewKV(client)
	kv.Put(context.TODO(), "/cron/jobs/job2", "job2 again")
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}

	log.Warn("2")
	if len(delResp.PrevKvs) != 0 {
		for index, value := range delResp.PrevKvs {
			fmt.Println("序号 =", index, " key= ", string(value.Key), " value= ", string(value.Value))
		}
	}

}

func main() {
	deleteValue()
}
