package utils

import (
	"bytes"
	"os/exec"
)

type CommandOutput struct {
	Output bytes.Buffer
	cmd    *exec.Cmd
}

func OutputCommand(name string, args ...string) *CommandOutput {
	cmdOut := CommandOutput{}
	cmdOut.cmd = exec.Command(name, args...)

	cmdOut.cmd.Stderr = &cmdOut.Output
	cmdOut.cmd.Stdout = &cmdOut.Output

	return &cmdOut
}

func (c *CommandOutput) Run() error {
	return c.cmd.Run()
}
