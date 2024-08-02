package elasticsearch

import (
	"dockit/constant"
	"dockit/util"
)

func Ports() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "9200",
		ports[1]: "9300",
	}
}

func Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/elasticsearch/config":  "/usr/share/elasticsearch/config",
		constant.Home + "/dockit/elasticsearch/logs":    "/usr/share/elasticsearch/logs",
		constant.Home + "/dockit/elasticsearch/plugins": "/usr/share/elasticsearch/plugins",
		constant.Home + "/dockit/elasticsearch/data":    "/usr/share/elasticsearch/data",
	}
}

func Environments() map[string]string {
	return map[string]string{
		"discovery.type": "single-node",
		"ES_JAVA_OPTS":   "-Xms512m -Xmx512m",
	}
}

func Others() []string {
	return []string{
		"--privileged",
	}
}
