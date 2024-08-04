package docker

import (
	"dockit/model"
	"dockit/util"
	"strings"
)

// Pull the mirror image.
func Pull(image *model.Image) {
	util.ExecCmd("docker", "pull", image.FullImageName())
	image.Id = getImageId(image)
}

// Run the image in a container.
func Run(container *model.Container, options []string) {
	// Generate the run args
	args := container.DefaultRunArgs()
	args = append(args, container.Option.GenerateCustomOptions()...)
	args = append(args, options...)
	args = append(args, container.Image.Id)
	util.ExecCmd("docker", args...)

	// Set the container ID
	container.Id = getContainerId(container.Name)
}

// Rm remove the container.
func Rm(image *model.Image) {
	util.ExecCmd("docker", "rm", "-f", model.GetContainerName(image))
}

// Rmi remove the image.
func Rmi(image *model.Image) {
	util.ExecCmd("docker", "rmi", image.FullImageName())
}

func getImageId(image *model.Image) string {
	output := util.ExecCmdWithOutput("docker", "image", "ls", image.FullImageName(), "-q")
	return strings.TrimSpace(output)
}

func getContainerId(containerName string) string {
	output := util.ExecCmdWithOutput("docker", "ps", "-a", "--filter", "name="+containerName+"", "--format", "{{.ID}}")
	return strings.TrimSpace(output)
}
