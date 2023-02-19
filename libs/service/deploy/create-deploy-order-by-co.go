package deploy_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateDeployOrderByCO(ctx context.Context, operator *dao.User, co *dao.CompileOrder, image string) (*dao.DeployOrder, error) {
	client := dbc.DB

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

	return deployDao, nil
}
