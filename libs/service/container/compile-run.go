package container_service

import (
	"errors"

	"github.com/sheason2019/spoved/exceptions/exception"
	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func CompileRun(dir string) *exception.Exception {
	// 检查build shell是否存在
	buildShellPath := dir + "/" + "build.sh"
	exist, err := file_service.Exist(buildShellPath)
	if err != nil {
		return exception.New(err)
	}
	if !exist {
		return exception.New(errors.New("项目下不存在build.sh文件"))
	}

	// 获取绝对路径
	dir, err = file_service.GetAbsPath(dir)
	if err != nil {
		return exception.New(err)
	}

	// 若存在则执行 build shell
	// exec.Command("sudo", "docker", "run", "--entrypoint")

	return nil
}
