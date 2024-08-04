package option

import (
	"dockit/model"
	"dockit/util"
)

var imageOptMap = make(map[string][]OptStrategy)

// Register a custom option.
func Register(image string, options ...OptStrategy) {
	imageOptMap[image] = options
}

// GetContainerOption Get the option by image name and version.
func GetContainerOption(image *model.Image) *model.Option {
	optStrategies := imageOptMap[image.Name]

	if len(optStrategies) == 0 {
		return nil
	} else if len(optStrategies) == 1 {
		return assemblyOption(optStrategies[0])
	}

	// Find the option by version.
	for i := len(optStrategies) - 1; i >= 0; i-- {
		if util.CompareVersion(image.Version, optStrategies[i].Version()) >= 0 {
			return assemblyOption(optStrategies[i])
		}
	}
	return assemblyOption(optStrategies[0])
}

func assemblyOption(opt OptStrategy) *model.Option {
	// Assembly option object pre-processing
	assemblyOptionHandler(opt)

	return &model.Option{
		Ports:        opt.Ports(),
		Volumes:      opt.Volumes(),
		Environments: opt.Environments(),
		Others:       opt.Others(),
	}
}

func assemblyOptionHandler(opt OptStrategy) {
	if len(opt.Volumes()) == 0 {
		return
	}

	// Mkdir the volumes dir
	for k := range opt.Volumes() {
		_ = util.Mkdir(k, opt.VolumesPerm())
	}
}
