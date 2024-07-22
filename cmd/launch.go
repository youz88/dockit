package cmd

import (
	"dockit/docker"
	"fmt"
	"strings"
)

func Launch(args []string) {
	if len(args) == 0 {
		missingParams()
	} else if len(args) != 1 || args[0] == "--help" {
		mainHelp()
	} else {
		imageName := buildImageName(args[0])
		docker.Pull(imageName)
		docker.Run(docker.GetImageId(imageName))
	}
}

// Build a mirror name
func buildImageName(imageName string) string {
	if !strings.Contains(imageName, ":") {
		imageName = imageName + ":latest"
	}
	return imageName
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
