package option

import (
	"dockit/constant"
	"dockit/util"
)

func init() {
	Register("nacos", Nacos{})
}

type Nacos struct {
	DefaultOpt
}

func (nacos Nacos) Ports() map[string]string {
	ports := util.GenPort(3)
	return map[string]string{
		ports[0]: "8848",
		ports[1]: "9848",
		ports[2]: "9849",
	}
}

func (nacos Nacos) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nacos/logs":                   "/home/nacos/logs",
		constant.Home + "/dockit/nacos/application.properties": "/home/nacos/conf/application.properties",
	}
}

func (nacos Nacos) Environments() map[string]string {
	return map[string]string{
		"MODE":             "standalone",
		"PREFER_HOST_MODE": "hostname",
	}
}

func (nacos Nacos) Others() []string {
	return []string{
		"--privileged",
	}
}
