package common

import (
	"dockit/util"
	"fmt"
)

func Help() {
	// 输出帮助信息
	util.PrintInfo("\nUsage: dockit COMMAND")
	util.PrintInfo("Commands:")
	util.PrintInfo(fmt.Sprintf("%-20s%-20s", "launch", "Pull the docker image and start the container"))
	util.PrintInfo(fmt.Sprintf("%-20s%-20s", "clear", "Clear the docker image and container"))
	util.PrintInfo("\nRun 'dockit COMMAND --help' for more information on a command.")
}
