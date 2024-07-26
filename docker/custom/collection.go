package custom

type Options interface {
	Ports() map[string]string
	Volumes() map[string]string
	Environments() map[string]string
	Others() []string
}

var ImageMap = make(map[string]Options)

func init() {
	ImageMap["redis"] = &Redis{}
	ImageMap["nacos"] = &Nacos{}
	ImageMap["nginx"] = &Nginx{}
	ImageMap["rabbitmq"] = &RabbitMQ{}
}
