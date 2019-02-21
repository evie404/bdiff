package changes

import "strings"

func IsBuildFile(file string) bool {
	return strings.HasSuffix(file, "BUILD") || strings.HasSuffix(file, "BUILD.bazel") || strings.HasSuffix(file, ".bzl")
}

func BuildSrcFiles(files []string) (buildFiles, srcFiles []string) {
	buildFiles = make([]string, 0, len(files))
	srcFiles = make([]string, 0, len(files))

	for _, file := range files {
		if IsBuildFile(file) {
			buildFiles = append(buildFiles, file)
		} else {
			srcFiles = append(srcFiles, file)
		}
	}

	return
}

func BuildFileChanges(file string) []string {
	if !IsBuildFile(file) {
		return nil
	}

	parts := strings.Split(file, "/")

	// TODO: rdeps(:all)
	if len(parts) <= 1 {
		return []string{"//..."}
	}

	// assume everything under the directory of build file changed for now
	return []string{"//" + strings.Join(parts[:len(parts)-1], "/") + "/..."}
}
