package docker

import (
	"dockit/docker/custom"
	"dockit/helper"
	"strings"
)

type Container struct {
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Image   *Image         `json:"image"`
	Options custom.Options `json:"options"`
}

// BuildContainerModel Build a container model.
func BuildContainerModel(image *Image) *Container {
	options := custom.ImageMap[image.Name]
	container := &Container{
		Name:    GetContainerName(image),
		Image:   image,
		Options: options,
	}
	return container
}

// GetCustomOptions Get custom options.
func (container *Container) GetCustomOptions() (options []string) {
	options = []string{
		"--name", GetContainerName(container.Image),
		"--restart=always",
	}
	if container.Options == nil {
		return options
	}

	options = append(options, buildOptions("-p", ":", container.Options.Ports())...)
	options = append(options, buildOptions("-v", ":", container.Options.Volumes())...)
	options = append(options, buildOptions("-e", "=", container.Options.Environments())...)
	options = append(options, container.Options.Others()...)
	return options
}

// RenderTable Render a container table output.
func (container *Container) RenderTable() {
	arr := [][]string{
		{
			"image-name", "image-id", "container-name", "container-id",
			"options",
		}, {
			container.Image.Name, container.Image.Id, container.Name, container.Id,
			strings.Join(container.GetCustomOptions(), " "),
		},
	}
	helper.RenderTable(arr)
}

// GetContainerName Get container name.
func GetContainerName(image *Image) string {
	return "dockit_" + image.Name
}

func buildOptions(prefix string, linkSymbol string, links map[string]string) (options []string) {
	if len(links) == 0 {
		return options
	}

	for key, val := range links {
		options = append(options, prefix, key+linkSymbol+val)
	}
	return options
}
