package filesystem

import "os"

func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	// may actually exist (if error is permission error) but we don't really care
	return false
}
