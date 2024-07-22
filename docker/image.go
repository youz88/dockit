package docker

import "strings"

type Image struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// FullImageName Get the full image name
func (image *Image) FullImageName() string {
	return image.Name + ":" + image.Version
}

// BuildImageModel Build a image model
func BuildImageModel(s string) *Image {
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
