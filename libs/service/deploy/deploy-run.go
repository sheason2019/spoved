package deploy_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 部署逻辑执行
func DeployRun(ctx context.Context, do *dao.DeployOrder) error {
	// 为DeployOrder注入关联信息
	err := dbc.DB.
		Preload("CompileOrder").
		Preload("CompileOrder.Project").
		Preload("CompileOrder.Project.Creator").
		Where("deploy_orders.id = ?", do.ID).
		Find(do).
		Error
	if err != nil {
		return err
	}

	// 创建Deployment
	err = k3s_service.Start(ctx, do)
	if err != nil {
		do.StatusCode = -1
	} else {
		do.StatusCode = 1
	}

	// 根据DeployOrder创建Service
	err = k3s_service.CreateServiceByDeployOrder(ctx, do)
	if err != nil {
		return err
	}

	// 由于网关服务没有与Spoved进行解耦，这里需要对部署Spoved的情况进行特殊处理
	// 在部署Spoved服务以后，需要调整ingress的入口
	proj := do.CompileOrder.Project
	if proj.Creator.Username == "root" && proj.ProjectName == "spoved" {
		_, err = k3s_service.UpdateSpovedIngress(ctx, do)
	}

	return err
}
