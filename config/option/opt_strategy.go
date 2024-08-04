package option

import (
	"dockit/constant"
	"io/fs"
)

type OptStrategy interface {
	Version() string
	Ports() map[string]string
	Volumes() map[string]string
	VolumesPerm() fs.FileMode
	Environments() map[string]string
	Others() []string
}

type DefaultOpt struct{}

func (d DefaultOpt) Version() string {
	return ""
}

func (d DefaultOpt) Ports() map[string]string {
	return nil
}

func (d DefaultOpt) Volumes() map[string]string {
	return nil
}

func (d DefaultOpt) VolumesPerm() fs.FileMode {
	return constant.DefaultFilePerm
}

func (d DefaultOpt) Environments() map[string]string {
	return nil
}

func (d DefaultOpt) Others() []string {
	return nil
}
