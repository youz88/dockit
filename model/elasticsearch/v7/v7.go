package v7

import (
	"dockit/constant"
	"dockit/model"
	"dockit/model/elasticsearch"
)

const Version7 = "7.0.0"

func init() {
	model.Register("elasticsearch", &model.Option{
		Version:      Version7,
		Ports:        elasticsearch.Ports(),
		Volumes:      volumes(),
		Environments: elasticsearch.Environments(),
		Others:       elasticsearch.Others(),
	})
}

func volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/elasticsearch/config":  "/usr/share/elasticsearch/config",
		constant.Home + "/dockit/elasticsearch/logs":    "/usr/share/elasticsearch/logs",
		constant.Home + "/dockit/elasticsearch/plugins": "/usr/share/elasticsearch/plugins",
		constant.Home + "/dockit/elasticsearch/data":    "/var/lib/elasticsearch/data",
	}
}
