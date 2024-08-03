package nginx

import (
	"dockit/constant"
	"dockit/docker"
	"dockit/model"
)

var defaultOption = &model.Option{
	Ports:        defaultPorts(),
	Volumes:      defaultVolumes(),
	Environments: defaultEnvironments(),
	Others:       defaultOthers(),
}

func init() {
	docker.Register("nginx", defaultOption)
}

func defaultPorts() map[string]string {
	return map[string]string{
		"80":  "80",
		"443": "443",
	}
}

func defaultVolumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nginx/conf/nginx.conf":   "/etc/nginx/nginx.conf",
		constant.Home + "/dockit/nginx/includ_conf":       "/etc/nginx/includ_conf",
		constant.Home + "/dockit/nginx/conf/default.conf": "/etc/nginx/conf.d/default.conf",
		constant.Home + "/dockit/nginx/html":              "/usr/share/nginx/html",
		constant.Home + "/dockit/nginx/logs":              "/var/log/nginx",
		constant.Home + "/dockit/nginx/certs":             "/etc/nginx/certs",
	}
}

func defaultEnvironments() map[string]string {
	return nil
}

func defaultOthers() []string {
	return []string{
		"--privileged",
	}
}
