package deploy_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindLatestDeployOrder(ctx context.Context, proj *dao.Project) (*dao.DeployOrder, error) {
	client := dbc.DB

	deploys := make([]dao.DeployOrder, 0)
	err := client.
		WithContext(ctx).
		Joins("CompileOrder", client.Where("CompileOrder.project_id = ?", proj.ID)).
		Limit(1).
		Order("deploy_orders.created_at desc").
		Find(&deploys).
		Error

	if err != nil {
		return nil, fmt.Errorf("FindLatestDeployOrderError:%w", err)
	}

	if len(deploys) == 0 {
		return nil, nil
	}

	return &deploys[0], nil
}
