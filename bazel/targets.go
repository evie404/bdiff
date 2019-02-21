package bazel

import (
	"strings"
)

func IsInternalTarget(target string) bool {
	return strings.Contains(target, "~")
}

func TargetsFromSrcs(dir, bazelBin string, srcFiles []string, debug bool) ([]string, string, error) {
	if srcFiles == nil || len(srcFiles) == 0 {
		return nil, "", nil
	}

	targets := make([]string, 0, len(srcFiles))

	rdeps := "rdeps(//..., set(" + strings.Join(srcFiles, " ") + "))"
	foundTargets, stderr, err := Query(dir, bazelBin, rdeps, debug)
	if err != nil {
		return nil, stderr, err
	}

	for _, foundTarget := range foundTargets {
		if IsInternalTarget(foundTarget) {
			continue
		}

		targets = append(targets, foundTarget)
	}

	return targets, "", nil
}
