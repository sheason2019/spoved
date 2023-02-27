package deploy_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 部署逻辑执行
func DeployRun(ctx context.Context, do *dao.DeployOrder) error {
	// 创建Deployment
	err := k3s_service.Start(ctx, do)
	if err != nil {
		do.StatusCode = -1
	} else {
		do.StatusCode = 1
	}

	// 根据DeployOrder创建Service
	return k3s_service.CreateServiceByDeployOrder(ctx, do)
}
