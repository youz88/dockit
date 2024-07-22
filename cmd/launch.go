package cmd

import (
	"dockit/common"
	"dockit/docker"
	"fmt"
)

func Launch(args []string) {
	if len(args) == 0 {
		missParams()
	} else if len(args) != 1 || args[0] == "--help" {
		mainHelp()
	} else {
		docker.Pull(args[0])
		common.ExecCmd("docker", "run", "-it", "-d", args[0])
	}
}

func mainHelp() {
	fmt.Println("\nUsage: dockit launch NAME[:TAG|@DIGEST]")
	fmt.Println("\nPull the docker image and start the container")
}

func missParams() {
	// 输出帮助信息
	fmt.Println("dockit launch requires exactly 1 argument.")
	fmt.Println("See 'dockit launch --help'.")
	mainHelp()
}
