package container_service

import (
	"errors"
	"os/exec"

	"github.com/sheason2019/spoved/exceptions/exception"
	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func CompileRun(dir string) (string, *exception.Exception) {
	// 检查build shell是否存在
	buildShellPath := dir + "/" + "build.sh"
	exist, err := file_service.Exist(buildShellPath)
	if err != nil {
		return "", exception.New(err)
	}
	if !exist {
		return "", exception.New(errors.New("项目下不存在build.sh文件"))
	}

	// 获取绝对路径
	dir, err = file_service.GetAbsPath(dir)
	if err != nil {
		return "", exception.New(err)
	}

	// 若存在则执行 build shell
	output, err := exec.Command(
		"sudo",
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
		return string(output), exception.New(err)
	}

	return string(output), nil
}
