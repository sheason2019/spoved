package deploy_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

// 寻找指定Project下正在执行的DeployOrder
func FindRunningDeployOrders(ctx context.Context, proj *dao.Project) ([]dao.DeployOrder, error) {
	client := dbc.DB

	deploys := make([]dao.DeployOrder, 0)
	err := client.
		WithContext(ctx).
		Joins("CompileOrder", client.Where("CompileOrder.project_id = ?", proj.ID)).
		Where("deploy_orders.service_name is not NULL").
		Where("deploy_orders.service_name is not ''").
		Find(&deploys).
		Error

	if err != nil {
		return nil, fmt.Errorf("FindLatestDeployOrderError:%w", err)
	}

	return deploys, nil
}
