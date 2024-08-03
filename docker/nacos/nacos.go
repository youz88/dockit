package nacos

import (
	"dockit/constant"
	"dockit/docker"
	"dockit/model"
	"dockit/util"
)

var defaultOption = &model.Option{
	Ports:        defaultPorts(),
	Volumes:      defaultVolumes(),
	Environments: defaultEnvironments(),
	Others:       defaultOthers(),
}

func init() {
	docker.Register("nacos", defaultOption)
}

func defaultPorts() map[string]string {
	ports := util.GenPort(3)
	return map[string]string{
		ports[0]: "8848",
		ports[1]: "9848",
		ports[2]: "9849",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nacos/logs":                   "/home/nacos/logs",
		constant.Home + "/dockit/nacos/application.properties": "/home/nacos/conf/application.properties",
	}
}

func defaultEnvironments() map[string]string {
	return map[string]string{
		"MODE":             "standalone",
		"PREFER_HOST_MODE": "hostname",
	}
}

func defaultOthers() []string {
	return []string{
		"--privileged",
	}
}
