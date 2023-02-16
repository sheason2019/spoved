package compile_service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	images "github.com/sheason2019/spoved/libs/service/images"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 编译项目
func CompileRun(ctx context.Context, co *dao.CompileOrder) (err error) {
	defer func(cause error) {
		if cause == nil {
			co.StatusCode = 1
		} else {
			co.StatusCode = -1
		}

		fmt.Println(err)

		dbc.DB.WithContext(ctx).Save(co)
	}(err)

	// TODO: 操作人权限校验

	// 镜像校验
	if !images.ValidateImage(co.Image, "compile") {
		return errors.WithStack(errors.New("不支持的镜像：" + co.Image))
	}

	// 拉取代码
	err = k3s_service.GitClone(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	// 执行编译命令
	err = k3s_service.Build(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
