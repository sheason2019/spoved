package deploy_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateDeployOrder(ctx context.Context, do *dao.DeployOrder) error {
	return dbc.DB.WithContext(ctx).Save(do).Error
}
