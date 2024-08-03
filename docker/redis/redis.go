package redis

import (
	"dockit/constant"
	"dockit/docker"
	"dockit/model"
	"dockit/util"
)

var defaultOption = &model.Option{
	Ports:   defaultPorts(),
	Volumes: defaultVolumes(),
}

func init() {
	docker.Register("redis", defaultOption)
}

func defaultPorts() map[string]string {
	ports := util.GenPort(1)
	return map[string]string{
		ports[0]: "6379",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/redis/conf/redis.conf": "/usr/local/etc/redis/redis.conf",
		constant.Home + "/dockit/redis/data":            "/data",
	}
}
