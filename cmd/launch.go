package cmd

import (
	"dockit/config/option"
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

func launchHandler(_ *cobra.Command, args []string) {
	// Pull image.
	image := model.NewImage(args[0])
	docker.Pull(image)

	// Run container.
	container := model.NewContainer(image, option.GetContainerOption(image))
	docker.Run(container, args[1:])

	// Render table.
	container.Render()
}
