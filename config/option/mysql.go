package option

import (
	"dockit/constant"
	"dockit/util"
)

func init() {
	Register("mysql", Mysql{})
}

type Mysql struct {
	DefaultOpt
}

func (mysql Mysql) Ports() map[string]string {
	return map[string]string{
		util.GenPort(): "3306",
	}
}

func (mysql Mysql) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/mysql/data": "/var/lib/mysql",
	}
}

func (mysql Mysql) Environments() map[string]string {
	return map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
	}
}
