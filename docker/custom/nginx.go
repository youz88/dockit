package custom

import "dockit/constant"

type Nginx struct{}

func (redis *Nginx) Ports() map[string]string {
	return map[string]string{
		"80":  "80",
		"443": "443",
	}
}

func (redis *Nginx) Volumes() map[string]string {
	return map[string]string{
		constant.Home + "/dockit/nginx/conf/nginx.conf":   "/etc/nginx/nginx.conf",
		constant.Home + "/dockit/nginx/includ_conf":       "/etc/nginx/includ_conf",
		constant.Home + "/dockit/nginx/conf/default.conf": "/etc/nginx/conf.d/default.conf",
		constant.Home + "/dockit/nginx/html":              "/usr/share/nginx/html",
		constant.Home + "/dockit/nginx/logs":              "/var/log/nginx",
		constant.Home + "/dockit/nginx/certs":             "/etc/nginx/certs",
	}
}

func (redis *Nginx) Environments() map[string]string {
	return nil
}

func (redis *Nginx) Others() []string {
	return []string{
		"--privileged=true",
	}
}
