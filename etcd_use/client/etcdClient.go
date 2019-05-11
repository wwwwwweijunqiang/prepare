package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// etcd客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"192.168.31.6:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("etcd连接没有出错")
	client = client

}
