package bazel

import "strings"

func IsInternalTarget(target string) bool {
	return strings.Contains(target, "~")
}
