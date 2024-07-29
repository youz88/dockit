package helper

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

const HelpPrintFormat = "  %-20s%-20s\n"

// MainHelp Output help information.
func MainHelp() {
	fmt.Println("\nUsage: dockit COMMAND")

	fmt.Println("Commands:")
	fmt.Printf(HelpPrintFormat, "launch", "Pull the docker image and start the container")
	fmt.Printf(HelpPrintFormat, "clear", "Clear the docker image and container")

	fmt.Println("\nRun 'dockit COMMAND --help' for more information on a command.")
}

// RenderTable Render table output.
func RenderTable(arr [][]string) {
	if len(arr) == 0 {
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(arr[0])
	table.AppendBulk(arr[1:])
	table.Render()
}

func printTableRowSeparator(columnSize int) {
	for i := 0; i < columnSize; i++ {
		fmt.Printf("+--------------------")
	}
	fmt.Printf("+\n")
}
