package common

import (
	"bytes"
	"dockit/util"
	"fmt"
	"os/exec"
)

var (
	stdout bytes.Buffer
	stderr bytes.Buffer
)

func ExecCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		util.PrintError(stderr.String())
		return err
	}

	// 打印输出
	fmt.Println(stdout.String())
	return nil
}
