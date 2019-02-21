package bazel

import (
	"log"
	"strings"

	"github.com/rickypai/bdiff/cmd"
)

func Query(dir, bazelBin, query string, debug bool) (targets []string, stderr string, err error) {
	if debug {
		log.Printf("$ bazel query \"%s\"\n", query)
	}

	out, stderr, err := cmd.ExecCommand(dir, bazelBin, "query", query)
	if err != nil {
		return nil, stderr, err
	}

	return strings.Split(string(out), "\n"), "", nil
}

func FileTracked(dir, bazelBin, file string, debug bool) (tracked bool, stderr string, err error) {
	targets, stderr, err := Query(dir, bazelBin, file, debug)
	if strings.Contains(stderr, "no such target") && strings.Contains(stderr, "however, a source file of this name exists.") {
		return false, stderr, nil
	}

	if err != nil {
		return false, stderr, err
	}

	if len(targets) > 0 {
		return true, "", nil
	}

	return false, "", nil
}

func FilesTracked(dir, bazelBin string, files []string, debug bool) (tracked, notTracked []string, stderr string, err error) {
	tracked = make([]string, 0, len(files))
	notTracked = make([]string, 0, len(files))

	for _, file := range files {
		isTracked, stderr, err := FileTracked(dir, bazelBin, file, debug)
		if err != nil {
			return nil, nil, stderr, err
		}

		if isTracked {
			tracked = append(tracked, file)
		} else {
			notTracked = append(notTracked, file)
		}
	}

	return
}
