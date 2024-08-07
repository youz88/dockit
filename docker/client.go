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

// GetContainerGateway Get the container gateway according to the container id.
func GetContainerGateway(containerId string) string {
	gateway := util.ExecCmdWithOutput("docker", "inspect",
		"-f", "{{range.NetworkSettings.Networks}}{{.Gateway}}{{end}}", containerId)
	return strings.TrimSpace(gateway)
}

// GetImageByName Get the image list according to the image name.
func GetImageByName(imageName string) []*model.Image {
	var images []*model.Image

	// Get the mirror list of docker images
	imageStr := util.ExecCmdWithOutput("docker", "images", "--format", "{{.ID}} {{.Repository}} {{.Tag}}", "--filter", "reference="+imageName+"*")
	if imageStr == "" {
		return images
	}

	// Split the mirror list string into an array by line break.
	imageStrArr := strings.Split(imageStr, "\n")
	for _, imageInfoStr := range imageStrArr {
		if imageInfoStr == "" {
			continue
		}

		// Split the mirror information string into an array by space.
		imageInfo := strings.Split(imageInfoStr, " ")
		image := &model.Image{
			Id:      imageInfo[0],
			Name:    imageInfo[1],
			Version: imageInfo[2],
		}
		images = append(images, image)
	}
	return images
}

// GetContainerByImageName Get a list of running containers according to the image name.
func GetContainerByImageName(imageName string) []*model.Container {
	var containers []*model.Container

	// Get the image list according to the image name
	images := GetImageByName(imageName)
	if 0 == len(images) {
		return containers
	}

	for _, image := range images {
		// Get a list of running containers.
		containerStr := util.ExecCmdWithOutput("docker", "ps", "--filter", "status=running",
			"--filter", "ancestor="+image.Id, "--format", "{{.ID}} {{.Names}} {{.Ports}}")
		if containerStr == "" {
			continue
		}

		// Split the container list string into an array by line break.
		containerStrArr := strings.Split(containerStr, "\n")
		for _, containerInfoStr := range containerStrArr {
			if containerInfoStr == "" {
				continue
			}

			// Split the container information string into an array by space.
			containerInfoArr := strings.Split(containerInfoStr, " ")
			container := &model.Container{
				Id:   containerInfoArr[0],
				Name: containerInfoArr[1],
				Option: &model.Option{
					Ports: parserPorts(containerInfoArr),
				},
				Image: image,
			}

			// Add the container to the list of results.
			containers = append(containers, container)
		}
	}
	return containers
}

func getImageId(image *model.Image) string {
	imageId := util.ExecCmdWithOutput("docker", "image", "ls", image.FullImageName(), "-q")
	return strings.TrimSpace(imageId)
}

func parserPorts(containerInfoArr []string) map[string]string {
	if len(containerInfoArr) < 3 {
		return nil
	}

	return util.ParserDockerPorts(containerInfoArr[2:])
}

func getContainerId(containerName string) string {
	output := util.ExecCmdWithOutput("docker", "ps", "-a", "--filter", "name="+containerName+"", "--format", "{{.ID}}")
	return strings.TrimSpace(output)
}
