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

	// Create a command.
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	// Run the command.
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

	// Get the standard output of the command.
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// Run the command.
	err := cmd.Run()
	if err != nil {
		errMsg := strings.ReplaceAll(strings.Split(stderrBuf.String(), "\n")[0], "docker", "dockit")
		fmt.Println(errMsg)
		os.Exit(2)
	}

	// The command is executed successfully.
	return stdoutBuf.String()
}
