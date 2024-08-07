package model

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type Container struct {
	Id     string
	Name   string
	Image  *Image
	Option *Option
}

// NewContainer Build a container model.
func NewContainer(image *Image, option *Option) *Container {
	return &Container{
		Name:   GetContainerName(image),
		Image:  image,
		Option: option,
	}
}

// GetContainerName Get container name.
func GetContainerName(image *Image) string {
	return "dockit_" + image.BaseName
}

// DefaultRunArgs Get docker run default args.
func (container *Container) DefaultRunArgs() []string {
	return []string{"run", "-d", "--name", container.Name}
}

// Render a container table output.
func (container *Container) Render() {
	arr := [][]string{
		{
			"image-name", "image-id", "container-name", "container-id",
			"options",
		}, {
			container.Image.FullImageName(), container.Image.Id, container.Name, container.Id,
			container.Option.FormatStr(),
		},
	}

	// Render container table.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetReflowDuringAutoWrap(false)
	table.SetAutoWrapText(false)
	table.SetHeader(arr[0])
	table.AppendBulk(arr[1:])
	table.Render()
}
