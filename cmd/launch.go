package cmd

import (
	"dockit/docker"
	"fmt"
)

func Launch(args []string) {
	if len(args) == 0 {
		missingParams()
	} else if len(args) != 1 || args[0] == "--help" {
		mainHelp()
	} else {
		// Pull image
		image := docker.BuildImageModel(args[0])
		docker.Pull(image)

		// Run container
		containerId, containerName := docker.Run(image)
		container := docker.BuildContainerModel(image, containerId, containerName)

		// Render table
		container.RenderTable()
	}
}

// Output help information
func mainHelp() {
	fmt.Println("\nUsage: dockit launch NAME[:TAG|@DIGEST]")
	fmt.Println("\nPull the docker image and start the container")
}

// Input parameters are missing
func missingParams() {
	fmt.Println("dockit launch requires exactly 1 argument.")
	fmt.Println("See 'dockit launch --help'.")
	mainHelp()
}
