package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/env"
	git_service "github.com/sheason2019/spoved/libs/service/git"
	images "github.com/sheason2019/spoved/libs/service/images"
)

// 编译项目
func CompileRun(ctx CompileContext, image, nextVersion, branch string, proj *dao.Project, username string) (outputs []string, err error) {
	defer func(cause error) {
		if cause == nil {
			ctx.CompileOrder.StatusCode = 1
		} else {
			ctx.CompileOrder.StatusCode = -1
		}

		dbc.GetClient().WithContext(ctx).Save(ctx.CompileOrder)
	}(err)

	// TODO: 操作人权限校验

	// 镜像校验
	if !images.ValidateImage(image, "compile") {
		return outputs, errors.WithStack(errors.New("不支持的镜像：" + image))
	}

	// Git代码复制到的地址
	projDir := env.DataRoot + proj.DirPath() + "/" + nextVersion

	// 拉取代码
	err = git_service.GitClone(proj.GitUrl, projDir, branch, username)
	if err != nil {
		return outputs, errors.WithStack(err)
	}

	// 执行编译命令
	err = CompileRunBuild(context.Background(), projDir)
	if err != nil {
		return outputs, errors.WithStack(err)
	}

	return outputs, nil
}
