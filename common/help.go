package common

import (
	"fmt"
)

const HelpPrintFormat = "  %-20s%-20s\n"

// MainHelp Output help information
func MainHelp() {
	fmt.Println("\nUsage: dockit COMMAND")

	fmt.Println("Commands:")
	fmt.Printf(HelpPrintFormat, "launch", "Pull the docker image and start the container")
	fmt.Printf(HelpPrintFormat, "clear", "Clear the docker image and container")

	fmt.Println("\nRun 'dockit COMMAND --help' for more information on a command.")
}
