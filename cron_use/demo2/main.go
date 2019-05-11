package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression

	next time.Time
}

func cron() {
	var (
		cronJob *CronJob
		expr    *cronexpr.Expression
		now     time.Time

		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)

	// 当前时间
	now = time.Now()

	// 创建任务1
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr: expr,
		next: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	// 创建任务2
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr: expr,
		next: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob

	// 启动调动协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		// 定时检查任务调度表
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.next.Before(now) || cronJob.next.Equal(now) {
					// 启动一个协程，执行这个任务
					go func(jobName string) {
						fmt.Println("执行：", jobName)
					}(jobName)
					// 计算下次调度时间
					cronJob.next = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间:", cronJob.next)
				}
			}
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}
		}
	}()

	time.Sleep(100 * time.Second)
}

func cronTest() {
	var (
		cronJob *CronJob
		expr    *cronexpr.Expression
		now     time.Time

		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)

	// 当前时间
	now = time.Now()

	// 创建任务1
	expr = cronexpr.MustParse("*/2 * * * * * *")
	cronJob = &CronJob{
		expr: expr,
		next: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	fmt.Println(cronJob.next)
	fmt.Println(now)

	if cronJob.next.Before(now) {
		fmt.Println("before")
	} else {
		fmt.Println("after")
	}

	time.Sleep(12 * time.Second)
}
func main() {
	cron()
	//cronTest()
}
