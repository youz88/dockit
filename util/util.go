package util

import (
	"strconv"
	"strings"
)

// CompareVersion 比较两个版本号字符串
func CompareVersion(v1, v2 string) int {
	v1Parts := strings.Split(v1, ".")
	v2Parts := strings.Split(v2, ".")

	for i := 0; i < max(len(v1Parts), len(v2Parts)); i++ {
		var v1Part, v2Part int
		if i < len(v1Parts) {
			v1Part, _ = strconv.Atoi(v1Parts[i])
		}
		if i < len(v2Parts) {
			v2Part, _ = strconv.Atoi(v2Parts[i])
		}

		if v1Part < v2Part {
			return -1
		} else if v1Part > v2Part {
			return 1
		}
	}

	return 0
}
