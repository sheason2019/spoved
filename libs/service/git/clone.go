package git_service

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/env"
	"github.com/sheason2019/spoved/libs/utils"
)

var path_root string

func init() {
	if env.IS_PRODUCT {
		path_root = "data"
	} else {
		path_root = utils.GetRootPath() + "/data"
	}
}

func GitClone(url, codeDir, branch, username string) (string, error) {
	codeDir = path_root + codeDir
	sshDir := path_root + "/account/" + username + "/.ssh"

	os.RemoveAll(codeDir)

	return gitClone(context.Background(), url, branch, codeDir, sshDir)
}

func gitClone(ctx context.Context, url, branch, codeDir, sshDir string) (string, error) {
	var err error
	var output string

	utils.TimeoutFunc(ctx, func(ctx context.Context) {
		os.RemoveAll(codeDir)
		cmd := utils.OutputCommand(
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
	}, 15*1000)

	if err != nil {
		return output, errors.WithStack(err)
	}
	return output, nil
}
