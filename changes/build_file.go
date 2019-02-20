package changes

import "strings"

func IsBuildFile(file string) bool {
	return strings.HasSuffix(file, "BUILD") || strings.HasSuffix(file, "BUILD.bazel") || strings.HasSuffix(file, ".bzl")
}

func BuildFileChanges(file string) []string {
	if !IsBuildFile(file) {
		return nil
	}

	parts := strings.Split(file, "/")

	if len(parts) <= 1 {
		return []string{"//..."}
	}

	// assume everything under the directory of build file changed for now
	return []string{"//" + strings.Join(parts[:len(parts)-1], "/") + "/..."}
}
