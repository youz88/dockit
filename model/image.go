package model

import "strings"

type Image struct {
	Id      string
	Name    string
	Version string
}

// FullImageName Get the full image name.
func (image *Image) FullImageName() string {
	return image.Name + ":" + image.Version
}

// NewImage Build an image model.
func NewImage(s string) *Image {
	var image Image
	if strings.Contains(s, ":") {
		split := strings.Split(s, ":")
		image.Name = split[0]
		image.Version = split[1]
	} else {
		image.Version = "latest"
		image.Name = s
	}
	return &image
}
