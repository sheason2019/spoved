package compile_service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	file_service "github.com/sheason2019/spoved/libs/service/file"
	"github.com/sheason2019/spoved/libs/utils"
	output_command "github.com/sheason2019/spoved/libs/utils/output-command"
)

func CompileRunBuild(ctx context.Context, dir string) (output string, err error) {
	// 检查build shell是否存在
	buildShellPath := dir + "/" + "build.sh"
	exist, err := file_service.Exist(buildShellPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if !exist {
		return "", errors.WithStack(errors.New("项目下不存在build.sh文件"))
	}

	// 执行编译逻辑，时限30分钟
	utils.TimeoutFunc(ctx, func(ctx context.Context, cancel func()) {
		cmd := output_command.Command(
			"docker",
			"run",
			"--entrypoint",
			"/bin/sh",
			"-v",
			dir+":/app",
			"node:16-alpine",
			"/app/build.sh",
		)

		err = cmd.Run()
		output = cmd.Output.String()
		cancel()
	}, 30*60*time.Second)

	return string(output), err
}
