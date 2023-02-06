package git_service

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/env"
	"github.com/sheason2019/spoved/libs/utils"
	output_command "github.com/sheason2019/spoved/libs/utils/output-command"
)

func GitClone(url, codeDir, branch, username string) (output string, err error) {
	sshDir := env.DataRoot + "/account/" + username + "/.ssh"

	utils.TimeoutFunc(context.Background(), func(ctx context.Context, cancel func()) {
		os.RemoveAll(codeDir)
		cmd := output_command.Command(
			"docker",
			"run",
			"-v",
			codeDir+":/code",
			"-v",
			sshDir+":/root/.ssh",
			"bitnami/git:latest",
			"/bin/bash",
			"-c",
			cloneSh(url, branch),
		)

		err = cmd.Run()
		output = cmd.Output.String()
		cancel()
	}, 15*time.Second)

	if err != nil {
		return output, errors.WithStack(err)
	}
	return output, nil
}
