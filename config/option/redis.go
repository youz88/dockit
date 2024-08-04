package option

import (
	"dockit/constant"
	"dockit/util"
)

func init() {
	Register("redis", Redis{})
}

type Redis struct {
	DefaultOpt
}

func (redis Redis) Ports() map[string]string {
	ports := util.GenPort(1)
	return map[string]string{
		ports[0]: "6379",
	}
}

func (redis Redis) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/redis/conf/redis.conf": "/usr/local/etc/redis/redis.conf",
		constant.Home + "/dockit/redis/data":            "/data",
	}
}
