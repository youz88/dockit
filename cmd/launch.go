package cmd

import (
	"dockit/docker"
	"dockit/model"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:                   "launch NAME[:TAG|@DIGEST]",
	Short:                 "Pull the docker image and start the container",
	DisableFlagsInUseLine: true,
	Run:                   launchHandler,
	Args:                  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(launchCmd)
}

func launchHandler(cmd *cobra.Command, args []string) {
	// Pull image.
	image := model.BuildImageModel(args[0])
	docker.Pull(image)

	// Run container.
	container := model.BuildContainerModel(image)
	docker.Run(container, args[1:])

	// Render table.
	container.Render()
}
