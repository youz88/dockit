package docker

import (
	"dockit/common"
	"strings"
)

// Pull the mirror image.
func Pull(image *Image) {
	common.ExecCmd("docker", "pull", image.FullImageName())
	image.Id = getImageId(image)
}

// Run the image in a container.
func Run(image *Image) (string, string) {
	containerName := "dockit_" + image.Name + ""
	common.ExecCmd("docker", "run", "-d", "--name", containerName, image.Id)
	return getContainerId(containerName), containerName
}

func getImageId(image *Image) string {
	output := common.ExecCmdWithOutput("docker", "image", "ls", image.FullImageName(), "-q")
	return strings.TrimSpace(output)
}

func getContainerId(containerName string) string {
	output := common.ExecCmdWithOutput("docker", "ps", "-a", "--filter", "name="+containerName+"", "--format", "{{.ID}}")
	return strings.TrimSpace(output)
}
