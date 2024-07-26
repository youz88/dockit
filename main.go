package main

import (
	"dockit/cmd"
	"dockit/helper"
	"os"
)

func main() {
	args := os.Args[1:]

	// Check if the command is valid.
	if len(args) == 0 || cmd.CommandMap[args[0]] == nil {
		helper.MainHelp()
		return
	}

	// Exec executes the command with the given arguments.
	command := cmd.CommandMap[args[0]]
	args = args[1:]
	if len(args) == 0 {
		command.MissingParams()
	} else if args[0] == "--help" {
		command.MainHelp()
	} else {
		command.CustomExec(args)
	}
}
