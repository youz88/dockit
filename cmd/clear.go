package cmd

import (
	"dockit/docker"
	"fmt"
)

type Clear struct{}

func (clear *Clear) MissingParams() {
	fmt.Println("dockit clear requires exactly 1 argument.")
	fmt.Println("See 'dockit clear --help'.")
}

func (clear *Clear) MainHelp() {
	fmt.Println("\nUsage: dockit clear NAME[:TAG|@DIGEST]")
	fmt.Println("\nRemove the docker image and container")
}

func (clear *Clear) CustomExec(args []string) {
	image := docker.BuildImageModel(args[0])
	docker.Rm(image)
	docker.Rmi(image)
}
