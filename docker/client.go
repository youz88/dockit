package docker

import "dockit/common"

func Pull(name string) {
	common.ExecCmd("docker", "pull", name)
}
