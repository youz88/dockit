package docker

import (
	"dockit/util"
	"strings"
)

// Pull the mirror image.
func Pull(image *Image) {
	util.ExecCmd("docker", "pull", image.FullImageName())
	image.Id = getImageId(image)
}

// Run the image in a container.
func Run(container *Container, options []string) {
	args := []string{"run", "-d"}
	args = append(args, container.GetCustomOptions()...)
	args = append(args, options...)
	args = append(args, container.Image.Id)
	util.ExecCmd("docker", args...)

	// Set the container ID
	container.Id = getContainerId(container.Name)
}

// Rm remove the container.
func Rm(image *Image) {
	util.ExecCmd("docker", "rm", "-f", GetContainerName(image))
}

// Rmi remove the image.
func Rmi(image *Image) {
	util.ExecCmd("docker", "rmi", image.FullImageName())
}

func getImageId(image *Image) string {
	output := util.ExecCmdWithOutput("docker", "image", "ls", image.FullImageName(), "-q")
	return strings.TrimSpace(output)
}

func getContainerId(containerName string) string {
	output := util.ExecCmdWithOutput("docker", "ps", "-a", "--filter", "name="+containerName+"", "--format", "{{.ID}}")
	return strings.TrimSpace(output)
}
