package container_service

import (
	"os/exec"

	"github.com/pkg/errors"
	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func CompileRun(dir string) (string, error) {
	// 检查build shell是否存在
	buildShellPath := dir + "/" + "build.sh"
	exist, err := file_service.Exist(buildShellPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if !exist {
		return "", errors.WithStack(errors.New("项目下不存在build.sh文件"))
	}

	// 获取绝对路径
	dir, err = file_service.GetAbsPath(dir)
	if err != nil {
		return "", errors.WithStack(err)
	}

	// 若存在则执行 build shell
	output, err := exec.Command(
		"docker",
		"run",
		"--entrypoint",
		"/bin/sh",
		"-v",
		dir+":/app",
		"node:16-alpine",
		"/app/build.sh",
	).Output()
	if err != nil {
		return string(output), errors.WithStack(err)
	}

	return string(output), nil
}
