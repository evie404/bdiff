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

	// ignore deleted files
	files, _ := filesystem.FilesExists(allFiles)

	// ignore untracked files
	trackedFiles, _, stderr, err := bazel.FilesTracked(dir, bazelBin, files, debug)
	if err != nil {
		println(stderr)
		log.Fatal(err)
	}

	buildFiles, srcFiles := changes.BuildSrcFiles(trackedFiles)

	var targets []string

	for _, file := range buildFiles {
		buildTargets := changes.BuildFileChanges(file)

		if len(buildTargets) > 0 {
			targets = append(targets, buildTargets...)
		}
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

	_, externalTargets := bazel.InternalTargets(finalTargets)

	for _, target := range externalTargets {
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
