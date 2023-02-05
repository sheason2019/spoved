// 文件服务，所有文件操作都需要使用这个service
package file_service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/sheason2019/spoved/exceptions/exception"
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

func Read(path string) (string, error) {
	c, e := os.ReadFile(path_root + path)
	return string(c), e
}

func Write(c, path string) error {
	exist, err := Exist(path)
	if err != nil {
		return err
	}

	path = path_root + path
	if !exist {
		os.MkdirAll(path[:strings.LastIndex(path, "/")], os.ModePerm)
	}

	return os.WriteFile(path, []byte(c), 0666)
}

func Exist(path string) (bool, error) {
	path = path_root + path

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func Mkdir(path string) {
	path = path_root + path

	os.MkdirAll(path, os.ModePerm)
}

func GitClone(url, dir, branch string) (string, *exception.Exception) {
	dir = path_root + dir

	fmt.Println(dir)

	outputs := []string{}

	os.RemoveAll(dir)

	e := gitClone(context.Background(), url, dir)
	if e != nil {
		return "", e.Wrap()
	}

	if branch != "master" {
		output, err := exec.Command(
			"sh",
			"-c",
			fmt.Sprintf("cd %s\ngit fetch origin %s\ngit checkout %s\ngit pull", dir, branch, branch),
		).Output()
		if err != nil {
			outputs = append(outputs, string(output))
		}
	}

	return strings.Join(outputs, "\n"), nil
}

func gitClone(ctx context.Context, url, dir string) *exception.Exception {
	toCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	cmdOut := struct {
		Output *bytes.Buffer
		err    error
	}{
		Output: bytes.NewBuffer([]byte{}),
	}

	go func(ctx context.Context) {
		os.RemoveAll(dir)
		cmd := exec.Command("git", "clone", "--progress", "--depth", "1", url, dir)
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
		return exception.New(errors.New("拉取仓库超时"))
	}

	if cmdOut.err != nil {
		return exception.New(cmdOut.err)
	}
	return nil
}

// func checkoutBranch(ctx context.Context, dir, branch string) {}

func GetAbsPath(p string) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}

	p = path_root + p

	return filepath.Abs(p)
}
