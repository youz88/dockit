package rabbitmq

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
}

func init() {
	docker.Register("rabbitmq", defaultOption)
}

func defaultPorts() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "5682",
		ports[1]: "15672",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/rabbitmq": "/var/lib/rabbitmq",
	}
}

func defaultEnvironments() map[string]string {
	return map[string]string{
		"RABBITMQ_DEFAULT_USER": "admin",
		"RABBITMQ_DEFAULT_PASS": "admin",
	}
}
