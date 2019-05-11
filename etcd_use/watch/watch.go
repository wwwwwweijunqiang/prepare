package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
	kv     clientv3.KV

	delResp *clientv3.DeleteResponse

	getResp *clientv3.GetResponse

	watchStartRevision int64

	watchs clientv3.Watcher

	watchRespChan <-chan clientv3.WatchResponse

	watchResp clientv3.WatchResponse

	event *clientv3.Event
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

func watch() {
	client = getConnect()

	kv = clientv3.NewKV(client)

	// 开启协程，模拟etcd中数据的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "this is job7")

			kv.Delete(context.TODO(), "/cron/jobs/job7")
			time.Sleep(time.Second)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	// 得到当前值
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}
	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值：", string(getResp.Kvs[0].Value))
	}

	// 当前etcd集群事务的id,修改操作会改变当前值
	watchStartRevision = getResp.Header.Revision + 1

	// 创建一个watchs
	watchs = clientv3.NewWatcher(client)

	// 启动监听

	fmt.Println("开始监听版本： ", watchStartRevision)
	watchRespChan = watchs.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				//fmt.Println("修改为 ", string(event.Kv.Value), " 原先值是 ", string(event.PrevKv.Value))
				fmt.Println("修改为 ", string(event.Kv.Value))
			case mvccpb.DELETE:
				fmt.Println("删除了 ", string(event.Kv.Value), " Revision: ", event.Kv.ModRevision)
			}
		}

	}
}

func main() {

	watch()
}
