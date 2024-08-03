package docker

import "dockit/model"

var ImageOptMap = make(map[string][]*model.Option)

// Register a custom option.
func Register(image string, options ...*model.Option) {
	ImageOptMap[image] = options
}
