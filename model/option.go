package model

import (
	"fmt"
	"strings"
)

type Volume string

type Option struct {
	Ports        map[string]string
	Volumes      map[string]string
	Environments map[string]string
	Others       []string
}

// GenerateCustomOptions generates custom options based on the provided container and optional options.
func (opt *Option) GenerateCustomOptions() (options []string) {
	options = []string{
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
func (opt *Option) FormatStr() string {
	var builder strings.Builder
	for i, option := range opt.GenerateCustomOptions() {
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
