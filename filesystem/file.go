package filesystem

import "os"

func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	// may actually exist (if error is permission error) but we don't really care
	return false
}

func FilesExists(files []string) (exist, notExist []string) {
	exist = make([]string, 0, len(files))
	notExist = make([]string, 0, len(files))

	for _, file := range files {
		if FileExists(file) {
			exist = append(exist, file)
		} else {
			notExist = append(notExist, file)
		}
	}

	return
}
