package deploy_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func Deploy(ctx context.Context, operator *dao.User, co *dao.CompileOrder, image string) (any, error) {
	client := dbc.GetClient()

	deployDao := &dao.DeployOrder{
		Image:        image,
		StatusCode:   0,
		CompileOrder: *co,
		Operator:     *operator,
	}
	// 创建部署工单
	err := client.WithContext(ctx).Save(deployDao).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 执行部署逻辑
	go DeployRun(ctx, deployDao)

	return deployDao, nil
}
