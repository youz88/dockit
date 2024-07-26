package cmd

type Command interface {
	MissingParams()
	MainHelp()
	CustomExec(args []string)
}

var CommandMap = make(map[string]Command)

func init() {
	CommandMap["launch"] = &Launch{}
	CommandMap["clear"] = &Clear{}
}
