package option

import (
	"dockit/constant"
	"dockit/util"
)

func init() {
	Register("rabbitmq", Rabbitmq{})
}

type Rabbitmq struct {
	DefaultOpt
}

func (rabbitmq Rabbitmq) Ports() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "5682",
		ports[1]: "15672",
	}
}

func (rabbitmq Rabbitmq) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/rabbitmq": "/var/lib/rabbitmq",
	}
}

func (rabbitmq Rabbitmq) Environments() map[string]string {
	return map[string]string{
		"RABBITMQ_DEFAULT_USER": "admin",
		"RABBITMQ_DEFAULT_PASS": "admin",
	}
}
