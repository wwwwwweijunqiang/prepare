package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {

	var (
		expe  *cronexpr.Expression
		error error

		now  time.Time
		next time.Time
	)

	// 每隔5分钟执行一次
	if expe, error = cronexpr.Parse("*/1 * * * *"); error != nil {
		fmt.Println(error)
		return
	}
	now = time.Now()
	next = expe.Next(now)
	fmt.Printf("当前时间= %s , 下次执行时间= %s", now, next)
	//fmt.Println()
	//
	//time.AfterFunc(next.Sub(now), func() {
	//	fmt.Println(time.Now())
	//})
	//
	//
	//
	//time.Sleep(2*time.Minute)

}
