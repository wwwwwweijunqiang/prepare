package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	var (
		config clientv3.Config

		client *clientv3.Client
		err    error
		kv     clientv3.KV

		putResp *clientv3.PutResponse
	)

	// 配置参数
	config = clientv3.Config{
		Endpoints:   []string{"192.168.31.6:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	// 读写etcd的键值对
	kv = clientv3.NewKV(client)
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello world2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("revision", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("preValue: ", string(putResp.PrevKv.Value))
		}

	}
}
