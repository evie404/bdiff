package cmd

import (
	"bytes"
	"os/exec"
)

func ExecCommand(dir string, name string, arg ...string) (string, string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return out.String(), stderr.String(), err
}
