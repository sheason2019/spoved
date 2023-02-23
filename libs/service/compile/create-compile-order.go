package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	images "github.com/sheason2019/spoved/libs/service/images"
)

func CreateCompileOrder(ctx context.Context, order *dao.CompileOrder) error {
	// 镜像校验
	if !images.ValidateImage(order.Image, "compile") {
		return errors.WithStack(errors.New("不支持的镜像：" + order.Image))
	}

	return dbc.DB.WithContext(ctx).Save(order).Error
}
