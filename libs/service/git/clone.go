package git_service

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

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

	outputs := []string{}

	os.RemoveAll(codeDir)

	err := gitClone(context.Background(), url, branch, codeDir, sshDir)
	if err != nil {
		return "", err
	}

	return strings.Join(outputs, "\n"), nil
}

func gitClone(ctx context.Context, url, branch, codeDir, sshDir string) error {
	toCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	cmdOut := struct {
		Output *bytes.Buffer
		err    error
	}{
		Output: bytes.NewBuffer([]byte{}),
	}

	go func(ctx context.Context) {
		os.RemoveAll(codeDir)
		cmd := exec.Command(
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
		fmt.Println(cmd.Args)
		cmd.Stderr = cmdOut.Output
		cmd.Stdout = cmdOut.Output

		err := cmd.Run()
		if err != nil {
			cmdOut.err = err
		}
		cancel()
	}(toCtx)

	select {
	case <-toCtx.Done():
		break
	case <-time.After(time.Second * 15):
		fmt.Println(cmdOut.Output.String())
		return errors.WithStack(errors.New("拉取仓库超时"))
	}

	if cmdOut.err != nil {
		fmt.Println(cmdOut.Output.String())
		return errors.WithStack(cmdOut.err)
	}
	return nil
}
