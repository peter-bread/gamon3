package runtime

import (
	"bytes"
	"io"
	"os/exec"
)

type Runner struct{}

// Run runs a command and returns stderr output and error exit code.
func (r Runner) Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	var stderr bytes.Buffer
	cmd.Stdout = io.Discard
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stderr.String(), err
}
