package compile_service

import (
	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dbc"
	images "github.com/sheason2019/spoved/libs/service/images"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 编译项目
func CompileRun(ctx CompileContext) (outputs []string, err error) {
	co := ctx.CompileOrder

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
	if !images.ValidateImage(co.Image, "compile") {
		return outputs, errors.WithStack(errors.New("不支持的镜像：" + co.Image))
	}

	// 拉取代码
	err = k3s_service.GitClone(ctx, co)
	if err != nil {
		return outputs, errors.WithStack(err)
	}

	// 执行编译命令
	err = k3s_service.Build(ctx, co)
	if err != nil {
		return outputs, errors.WithStack(err)
	}

	return outputs, nil
}
