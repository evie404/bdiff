package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rickypai/bdiff/bazel"
	"github.com/rickypai/bdiff/changes"
	"github.com/rickypai/bdiff/cmd"
	"github.com/rickypai/bdiff/filesystem"
)

var (
	debugFlag     = flag.Bool("debug", false, "debug")
	bazelBinFlag  = flag.String("bazel-bin", "bazel", "bazel-bin")
	baseRefFlag   = flag.String("base", "", "base ref")
	targetRefFlag = flag.String("target", "HEAD", "target ref")
	testsOnlyFlag = flag.Bool("tests-only", false, "debug")
)

func main() {
	bazelBin, baseRef, targetRef, debug, testsOnly := parseFlags()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Println(dir)
	}

	// TODO: check workspace status and use that if dirty
	out, stderr, err := cmd.ExecCommand(dir, "git", "diff", "--name-only", fmt.Sprintf("%s..%s", baseRef, targetRef))
	if err != nil {
		println(out)
		println(stderr)
		log.Fatal(err)
	}

	allFiles := strings.Split(string(out), "\n")

	// assume everything changed if WORKSPACE changed
	for _, file := range allFiles {
		if file == "WORKSPACE" {
			fmt.Println("//...")

			return
		}
	}

	srcFiles := make([]string, 0, len(allFiles))
	delFiles := make([]string, 0, len(allFiles))

	var targets []string

	for _, file := range allFiles {
		if debug {
			println(file)
		}

		if len(file) < 1 {
			continue
		}

		// ignore deleted files for now
		if !filesystem.FileExists(file) {
			delFiles = append(delFiles, file)
			continue
		}

		buildTargets := changes.BuildFileChanges(file)

		if len(buildTargets) > 0 {
			targets = append(targets, buildTargets...)
			continue
		}

		// check if file is tracked by Bazel
		tracked, stderr, err := bazel.FileTracked(dir, bazelBin, file, debug)
		if err != nil {
			println(stderr)
			log.Fatal(err)
		}

		// don't bother with files not tracked by Bazel
		if !tracked {
			continue
		}

		srcFiles = append(srcFiles, file)
	}

	srcTargets, stderr, err := bazel.TargetsFromSrcs(dir, bazelBin, srcFiles, debug)
	if err != nil {
		println(stderr)
		log.Fatal(err)
	}

	targets = append(targets, srcTargets...)

	var set string

	if testsOnly {
		set = "tests(set(" + strings.Join(targets, " ") + "))"
	} else {
		set = "set(" + strings.Join(targets, " ") + ")"
	}

	finalTargets, stderr, err := bazel.Query(dir, bazelBin, set, debug)
	if err != nil {
		println(stderr)
		log.Fatal(err)
	}

	for _, target := range finalTargets {
		fmt.Println(target)
	}
}

func parseFlags() (bazelBin, baseRef, targetRef string, debug, testsOnly bool) {
	flag.Parse()

	if bazelBinFlag != nil {
		bazelBin = *bazelBinFlag
	} else {
		bazelBin = "bazelisk"
	}

	if *debugFlag {
		debug = true
	}

	if baseRefFlag != nil {
		baseRef = *baseRefFlag
	} else {
		baseRef = ""
	}

	if targetRefFlag != nil {
		targetRef = *targetRefFlag
	} else {
		targetRef = ""
	}

	if *testsOnlyFlag {
		testsOnly = true
	}

	return
}
