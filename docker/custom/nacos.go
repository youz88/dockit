package custom

import "dockit/constant"

type Nacos struct{}

func (redis *Nacos) Ports() map[string]string {
	return map[string]string{
		"8848": "8848",
		"9848": "9848",
		"9849": "9849",
	}
}

func (redis *Nacos) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nacos/logs":                   "/home/nacos/logs",
		constant.Home + "/dockit/nacos/application.properties": "/home/nacos/conf/application.properties",
	}
}

func (redis *Nacos) Environments() map[string]string {
	return map[string]string{
		"MODE":             "standalone",
		"PREFER_HOST_MODE": "hostname",
	}
}

func (redis *Nacos) Others() []string {
	return []string{
		"--privileged=true",
	}
}
