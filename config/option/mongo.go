package option

import (
	"dockit/constant"
	"dockit/util"
)

func init() {
	Register("mongo", Mongo{})
}

type Mongo struct {
	DefaultOpt
}

func (mongo Mongo) Ports() map[string]string {
	return map[string]string{
		util.GenPort(): "27017",
	}
}

func (mongo Mongo) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/mongodb/data": "/data/db",
	}
}
