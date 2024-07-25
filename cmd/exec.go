package cmd

type Command interface {
	MissingParams()
	MainHelp()
	CustomExec(args []string)
}

// Exec executes the command with the given arguments.
func Exec(cmd Command, args []string) {
	if len(args) == 0 {
		cmd.MissingParams()
	} else if args[0] == "--help" {
		cmd.MainHelp()
	} else {
		cmd.CustomExec(args)
	}
}
