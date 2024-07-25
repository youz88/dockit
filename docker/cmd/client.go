package docker

import (
	docker "dockit/docker/model"
	"dockit/util"
	"strings"
)

// Pull the mirror image.
func Pull(image *docker.Image) {
	util.ExecCmd("docker", "pull", image.FullImageName())
	image.Id = getImageId(image)
}

// Run the image in a container.
func Run(image *docker.Image, options []string) (string, string) {
	containerName := docker.BuildContainerName(image)
	args := []string{"run", "-d", "--name", containerName}
	args = append(args, options...)
	args = append(args, image.Id)
	util.ExecCmd("docker", args...)
	return getContainerId(containerName), containerName
}

// Rm remove the container.
func Rm(image *docker.Image) {
	util.ExecCmd("docker", "rm", "-f", docker.BuildContainerName(image))
}

// Rmi remove the image.
func Rmi(image *docker.Image) {
	util.ExecCmd("docker", "rmi", image.FullImageName())
}

func getImageId(image *docker.Image) string {
	output := util.ExecCmdWithOutput("docker", "image", "ls", image.FullImageName(), "-q")
	return strings.TrimSpace(output)
}

func getContainerId(containerName string) string {
	output := util.ExecCmdWithOutput("docker", "ps", "-a", "--filter", "name="+containerName+"", "--format", "{{.ID}}")
	return strings.TrimSpace(output)
}
