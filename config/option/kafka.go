package option

import (
	"dockit/config/exception"
	"dockit/constant"
	"dockit/docker"
	"dockit/util"
	"strings"
)

func init() {
	Register("kafka", Kafka{})
}

type Kafka struct {
	DefaultOpt
}

func (kafka Kafka) Ports() map[string]string {
	return map[string]string{
		util.GenPort(): "9092",
	}
}

func (kafka Kafka) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/kafka/data": "/var/lib/kafka/data",
	}
}

func (kafka Kafka) Environments() map[string]string {
	return map[string]string{
		"KAFKA_BROKER_ID":                        "1",
		"KAFKA_ZOOKEEPER_CONNECT":                getZookeeperConnect(),
		"KAFKA_ADVERTISED_LISTENERS":             "PLAINTEXT://localhost:9092",
		"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR": "1",
	}
}

func getZookeeperConnect() string {
	containers := docker.GetContainerByImageName(ZookeeperName)
	if 0 == len(containers) {
		exception.ZookeeperConnectError()
	}

	container := containers[0]
	for _, c := range containers {
		if strings.HasPrefix(c.Name, "dockit") {
			container = c
		}
	}

	gateway := docker.GetContainerGateway(container.Id)
	for k, v := range container.Option.Ports {
		if v == "2181" {
			return gateway + ":" + k
		}
	}

	exception.ZookeeperConnectError()
	return ""
}
