package redis

import (
	"dockit/constant"
	"dockit/model"
	"dockit/util"
)

func init() {
	model.Register("redis", &Redis{})
}

type Redis struct{}

func (opt *Redis) Ports() map[string]string {
	ports := util.GenPort(1)
	return map[string]string{
		ports[0]: "6379",
	}
}

func (opt *Redis) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/redis/conf/redis.conf": "/usr/local/etc/redis/redis.conf",
		constant.Home + "/dockit/redis/data":            "/data",
	}
}

func (opt *Redis) Environments() map[string]string {
	return nil
}

func (opt *Redis) Others() []string {
	return nil
}
