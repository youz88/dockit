package main

import (
	"dockit/cmd"
	"dockit/common"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		common.MainHelp()
		return
	}

	if args[0] == "launch" {
		cmd.Launch(args[1:])
	} else {
		common.MainHelp()
	}
}
