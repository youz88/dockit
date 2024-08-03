package elasticsearch

import (
	"dockit/constant"
	"dockit/docker"
	"dockit/model"
	"dockit/util"
)

var defaultOption = &model.Option{
	Ports:        defaultPorts(),
	Volumes:      defaultVolumes(),
	Environments: defaultEnvironments(),
	Others:       defaultOthers(),
}

func init() {
	docker.Register("elasticsearch", defaultOption)
}

func defaultPorts() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "9200",
		ports[1]: "9300",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/elasticsearch/plugins": "/usr/share/elasticsearch/plugins",
		constant.Home + "/dockit/elasticsearch/data":    "/usr/share/elasticsearch/data",
	}
}

func defaultEnvironments() map[string]string {
	return map[string]string{
		"discovery.type": "single-node",
		"ES_JAVA_OPTS":   "-Xms512m -Xmx512m",
	}
}

func defaultOthers() []string {
	return []string{
		"--privileged",
	}
}
