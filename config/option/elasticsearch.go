package option

import (
	"dockit/constant"
	"dockit/util"
	"io/fs"
)

func init() {
	Register("elasticsearch", Elasticsearch{})
}

type Elasticsearch struct {
	DefaultOpt
}

func (es Elasticsearch) Ports() map[string]string {
	ports := util.GenPort(2)
	return map[string]string{
		ports[0]: "9200",
		ports[1]: "9300",
	}
}

func (es Elasticsearch) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/elasticsearch/logs":    "/usr/share/elasticsearch/logs",
		constant.Home + "/dockit/elasticsearch/plugins": "/usr/share/elasticsearch/plugins",
		constant.Home + "/dockit/elasticsearch/data":    "/usr/share/elasticsearch/data",
	}
}

func (es Elasticsearch) VolumesPerm() fs.FileMode {
	return 0775
}

func (es Elasticsearch) Environments() map[string]string {
	return map[string]string{
		"discovery.type": "single-node",
		"ES_JAVA_OPTS":   "-Xms512m -Xmx512m",
	}
}

func (es Elasticsearch) Others() []string {
	return []string{
		"--privileged",
	}
}
