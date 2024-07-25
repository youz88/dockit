package model

import "dockit/helper"

type Container struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image *Image `json:"image"`
}

// BuildContainerName Build a container name.
func BuildContainerName(image *Image) string {
	return "dockit_" + image.Name
}

// BuildContainerModel Build a container model.
func BuildContainerModel(image *Image, containerId string, containerName string) *Container {
	return &Container{
		Id:    containerId,
		Name:  containerName,
		Image: image,
	}
}

// RenderTable Render a container table output.
func (container *Container) RenderTable() {
	arr := [][]string{
		{
			"image-name", "image-id", "container-name", "container-id",
		}, {
			container.Image.Name, container.Image.Id, container.Name, container.Id,
		},
	}
	helper.RenderTable(arr)
}
