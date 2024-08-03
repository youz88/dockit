package main

import (
	"dockit/cmd"
	_ "dockit/docker/elasticsearch"
	_ "dockit/docker/mysql"
	_ "dockit/docker/nacos"
	_ "dockit/docker/nginx"
	_ "dockit/docker/rabbitmq"
	_ "dockit/docker/redis"
)

func main() {
	cmd.Execute()
}
