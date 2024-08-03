package model

import (
	"dockit/util"
	"fmt"
	"strings"
)

type Option struct {
	Version      string
	Ports        map[string]string
	Volumes      map[string]string
	Environments map[string]string
	Others       []string
}

// GetImageOption get option by image.
func GetImageOption(image *Image, options []*Option) *Option {
	if len(options) == 0 {
		return nil
	} else if len(options) == 1 {
		return options[0]
	}

	for i := len(options) - 1; i >= 0; i-- {
		if util.CompareVersion(image.Version, options[i].Version) >= 0 {
			return options[i]
		}
	}
	return options[0]
}

// GenerateCustomOptions generates custom options based on the provided container and optional options.
func (opt *Option) GenerateCustomOptions(container *Container) (options []string) {
	options = []string{
		"--name", GetContainerName(container.Image),
		"--restart=always",
	}
	if opt == nil {
		return options
	}

	options = append(options, buildOptions("-p", ":", opt.Ports)...)
	options = append(options, buildOptions("-v", ":", opt.Volumes)...)
	options = append(options, buildOptions("-e", "=", opt.Environments)...)
	options = append(options, opt.Others...)
	return options
}

// FormatStr formats the custom options to a string.
func (opt *Option) FormatStr(container *Container) string {
	var builder strings.Builder
	for i, option := range opt.GenerateCustomOptions(container) {
		if i == 0 {
			builder.WriteString(fmt.Sprintf("%s ", option))
		} else if strings.HasPrefix(option, "-") {
			builder.WriteString(fmt.Sprintf(" \\\n%s ", option))
		} else {
			builder.WriteString(option)
		}
	}
	return builder.String()
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
