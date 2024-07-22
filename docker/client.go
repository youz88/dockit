package docker

import "dockit/common"

func Pull(imageName string) {
	common.ExecCmd("docker", "pull", imageName)
}

func Run(imageId string) {
	common.ExecCmd("docker", "run", "-it", "-d", imageId)
}

func GetImageId(imageName string) string {
	return common.ExecCmdWithOutput("docker", "image", "ls", imageName, "-q")
}
