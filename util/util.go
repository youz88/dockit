package util

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func Mkdir(path string, auth fs.FileMode) error {
	if len(path) == 0 {
		return fmt.Errorf("file path is empty")
	}

	// Determine whether the path type is a folder.
	splitPath := strings.Split(path, "/")
	if !strings.Contains(splitPath[len(splitPath)-1], ".") {
		path = path + "/"
	}

	// Create a folder.
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)

	// Modify permissions.
	if auth != 0 {
		_ = os.Chmod(dir, auth)
	}
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}
