package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "curl http://man.linuxde.net/curl")
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(string(output))
}
