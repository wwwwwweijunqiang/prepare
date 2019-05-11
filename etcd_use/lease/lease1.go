package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	getResp *clientv3.GetResponse

	lease clientv3.Lease

	leaseGrantResp *clientv3.LeaseGrantResponse

	id clientv3.LeaseID

	putResp *clientv3.PutResponse
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

func leases() {

	client = getConnect()

	// 创建一个租约
	lease = clientv3.NewLease(client)

	// 设置租约时间   10s
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	// 拿到租约id
	id = leaseGrantResp.ID

	var keepResp *clientv3.LeaseKeepAliveResponse

	var keepRespChan <-chan *clientv3.LeaseKeepAliveResponse

	// 自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), id); err != nil {
		fmt.Println(err)
		return
	}
	// 处理续租应答的协程
	go func() {
		for {

			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {

					fmt.Println("续租已经失效")
					goto END
				} else {
					fmt.Println("收到自动续租应答：", keepResp.ID)
				}

			}
		}
	END:
	}()

	// 得到Kv对象
	kv = clientv3.NewKV(client)

	// put数据，关联租约
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/lock1", "lock1", clientv3.WithLease(id)); err != nil {
		fmt.Println(err)
		return
	}

	log.Warn("写入成功", putResp.Header.Revision)

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/lock1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			log.Warn("租约过期")
			break
		}
		log.Warn("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}

func main() {
	leases()
}
