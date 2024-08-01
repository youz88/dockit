package cmd

import (
	"dockit/docker"
	"dockit/model"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear NAME[:TAG|@DIGEST]",
	Short: "Remove the docker image and container",
	Run:   clearHandler,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func clearHandler(cmd *cobra.Command, args []string) {
	image := model.BuildImageModel(args[0])
	docker.Rm(image)
	docker.Rmi(image)
}
