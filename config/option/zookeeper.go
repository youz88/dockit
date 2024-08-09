package option

import (
	"dockit/constant"
	"dockit/util"
)

const ZookeeperName = "zookeeper"

func init() {
	Register(ZookeeperName, Zookeeper{})
}

type Zookeeper struct {
	DefaultOpt
}

func (zk Zookeeper) Ports() map[string]string {
	return map[string]string{
		util.GenPort(): "2181",
	}
}

func (zk Zookeeper) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/zookeeper/data": "/data",
	}
}
