package cmd

import (
	docker "dockit/docker/cmd"
	"dockit/docker/model"
)

type Clear struct{}

func (clear *Clear) MissingParams() {
}

func (clear *Clear) MainHelp() {
}

func (clear *Clear) CustomExec(args []string) {
	image := model.BuildImageModel(args[0])
	docker.Rm(image)
	docker.Rmi(image)
}
