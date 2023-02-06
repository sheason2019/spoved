package output_command

import (
	"bytes"
	"os/exec"
)

type CommandOutput struct {
	Output bytes.Buffer
	Cmd    *exec.Cmd
}

func New() *CommandOutput {
	return &CommandOutput{}
}
