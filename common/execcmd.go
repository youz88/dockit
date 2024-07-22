package common

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecCmd Execute the command and output the standard output and standard errors directly to the console.
func ExecCmd(name string, args ...string) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		errMsg := strings.ReplaceAll(strings.Split(stderrBuf.String(), "\n")[0], "docker", "dockit")
		fmt.Println(errMsg)
		os.Exit(1)
		return
	}
}

// ExecCmdWithOutput Command, and return standard output and possible errors.
func ExecCmdWithOutput(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	// 获取命令的标准输出
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// 运行命令
	err := cmd.Run()
	if err != nil {
		// 命令执行失败，打印错误信息（包括标准错误输出）
		fmt.Println(fmt.Sprint(err) + ": " + stderrBuf.String())
		return ""
	}

	// 命令执行成功，打印标准输出
	return stdoutBuf.String()
}
