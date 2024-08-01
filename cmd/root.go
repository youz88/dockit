package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "dockit",
	Long: "Dockit is a Docker CLI package with built-in defaults for faster container ops.",
	Args: cobra.MinimumNArgs(2),
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
