package docker

import "dockit/common"

type Container struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image *Image `json:"image"`
}

func BuildContainerModel(image *Image, containerId string, containerName string) *Container {
	return &Container{
		Id:    containerId,
		Name:  containerName,
		Image: image,
	}
}

func (container *Container) RenderTable() {
	arr := [][]string{
		{
			"image-name", "image-id", "container-name", "container-id",
		}, {
			container.Image.Name, container.Image.Id, container.Name, container.Id,
		},
	}
	common.RenderTable(arr)
}
