package main

import (
	"dockit/cmd"
	"dockit/helper"
	"os"
)

var cmdMap = make(map[string]cmd.Command)

func init() {
	cmdMap["launch"] = &cmd.Launch{}
	cmdMap["clear"] = &cmd.Clear{}
}

func main() {
	args := os.Args[1:]

	// Check if the command is valid.
	if len(args) == 0 || cmdMap[args[0]] == nil {
		helper.MainHelp()
		return
	}

	// Execute the command.
	cmd.Exec(cmdMap[args[0]], args[1:])
}
