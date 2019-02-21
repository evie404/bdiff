package filesystem

import (
	"os"
	"strings"
)

func IsDir(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func DirOfFile(file string) string {
	if IsDir(file) {
		return file
	}

	parts := strings.Split(file, "/")
	if len(parts) < 2 {
		return ""
	}

	return strings.Join(parts[:len(parts)-1], "/")
}
