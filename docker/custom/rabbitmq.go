package custom

import "dockit/constant"

type RabbitMQ struct{}

func (redis *RabbitMQ) Ports() map[string]string {
	return map[string]string{
		"5682":  "5682",
		"15672": "15672",
	}
}

func (redis *RabbitMQ) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/rabbitmq": "/var/lib/rabbitmq",
	}
}

func (redis *RabbitMQ) Environments() map[string]string {
	return map[string]string{
		"RABBITMQ_DEFAULT_USER": "admin",
		"RABBITMQ_DEFAULT_PASS": "admin",
	}
}

func (redis *RabbitMQ) Others() []string {
	return nil
}
