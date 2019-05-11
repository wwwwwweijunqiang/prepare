package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {

	var (
		ctx        context.Context
		cancel     context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result

		res *result
	)

	resultChan = make(chan *result, 1000)

	ctx, cancel = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "echo hello;dir;echo world;")

		// 执行任务，捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出到main协程
		resultChan <- &result{
			err,
			output,
		}
	}()

	time.Sleep(1 * time.Second)

	// 取消上下文
	cancel()

	res = <-resultChan

	fmt.Println(res.err, string(res.output))
}
