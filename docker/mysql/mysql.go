package mysql

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
	docker.Register("mysql", defaultOption)
}

func defaultPorts() map[string]string {
	ports := util.GenPort(1)
	return map[string]string{
		ports[0]: "3306",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/mysql/data": "/var/lib/mysql",
	}
}

func defaultEnvironments() map[string]string {
	return map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
	}
}
