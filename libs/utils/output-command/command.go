package output_command

import (
	"os/exec"
)

func Command(name string, args ...string) *CommandOutput {
	cmdOut := New()
	cmdOut.Cmd = exec.Command(name, args...)

	cmdOut.Cmd.Stderr = &cmdOut.Output
	cmdOut.Cmd.Stdout = &cmdOut.Output

	return cmdOut
}

func (c *CommandOutput) Run() error {
	return c.Cmd.Run()
}
