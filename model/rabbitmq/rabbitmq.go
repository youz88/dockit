package rabbitmq

import (
	"dockit/constant"
	"dockit/model"
	"dockit/util"
)

func init() {
	model.Register("rabbitmq", &RabbitMQ{})
}

type RabbitMQ struct{}

func (opt *RabbitMQ) Ports() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "5682",
		ports[1]: "15672",
	}
}

func (opt *RabbitMQ) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/rabbitmq": "/var/lib/rabbitmq",
	}
}

func (opt *RabbitMQ) Environments() map[string]string {
	return map[string]string{
		"RABBITMQ_DEFAULT_USER": "admin",
		"RABBITMQ_DEFAULT_PASS": "admin",
	}
}

func (opt *RabbitMQ) Others() []string {
	return nil
}
