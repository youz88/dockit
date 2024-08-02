package nacos

import (
	"dockit/constant"
	"dockit/model"
	"dockit/util"
)

func init() {
	model.Register("nacos", &Nacos{})
}

type Nacos struct{}

func (opt *Nacos) Ports() map[string]string {
	ports := util.GenPort(3)
	return map[string]string{
		ports[0]: "8848",
		ports[1]: "9848",
		ports[2]: "9849",
	}
}

func (opt *Nacos) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nacos/logs":                   "/home/nacos/logs",
		constant.Home + "/dockit/nacos/application.properties": "/home/nacos/conf/application.properties",
	}
}

func (opt *Nacos) Environments() map[string]string {
	return map[string]string{
		"MODE":             "standalone",
		"PREFER_HOST_MODE": "hostname",
	}
}

func (opt *Nacos) Others() []string {
	return []string{
		"--privileged",
	}
}
