package v8

import (
	"dockit/model"
	"dockit/model/elasticsearch"
)

const Version8 = "8.0.0"

func init() {
	model.Register("elasticsearch", &model.Option{
		Version:      Version8,
		Ports:        elasticsearch.Ports(),
		Volumes:      elasticsearch.Volumes(),
		Environments: elasticsearch.Environments(),
		Others:       elasticsearch.Others(),
	})
}
