package custom

type Nginx struct{}

func (redis *Nginx) Ports() map[string]string {
	return map[string]string{
		"80":  "80",
		"443": "443",
	}
}

func (redis *Nginx) Volumes() map[string]string {
	return map[string]string{
		"~/dockit/nginx/conf/nginx.conf":   "/etc/nginx/nginx.conf",
		"~/dockit/nginx/includ_conf":       "/etc/nginx/includ_conf",
		"~/dockit/nginx/conf/default.conf": "/etc/nginx/conf.d/default.conf",
		"~/dockit/nginx/html":              "/usr/share/nginx/html",
		"~/dockit/nginx/logs":              "/var/log/nginx",
		"~/dockit/nginx/certs":             "/etc/nginx/certs",
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
