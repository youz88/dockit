package mysql

import (
	"dockit/constant"
	"dockit/model"
	"dockit/util"
)

func init() {
	model.Register("mysql", &Mysql{})
}

type Mysql struct{}

func (opt *Mysql) Ports() map[string]string {
	ports := util.GenPort(1)
	return map[string]string{
		ports[0]: "3306",
	}
}

func (opt *Mysql) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/mysql/data": "/var/lib/mysql",
	}
}

func (opt *Mysql) Environments() map[string]string {
	return map[string]string{
		"MYSQL_ROOT_PASSWORD": "root",
	}
}

func (opt *Mysql) Others() []string {
	return nil
}
