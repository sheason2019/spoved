package deploy_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 部署逻辑执行
func DeployRun(ctx context.Context, do *dao.DeployOrder) (err error) {
	// 结束部署逻辑后修改DeployOrder的状态
	defer func() {
		if err != nil {
			do.StatusCode = -1
		} else {
			do.StatusCode = 1
		}
		dbc.DB.Save(do)
	}()
	// 为DeployOrder注入关联信息
	err = dbc.DB.
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
	err = k3s_service.CreateDeploymentByDeployOrder(ctx, do)
	if err != nil {
		return err
	}

	// 根据DeployOrder创建Service
	err = k3s_service.CreateServiceByDeployOrder(ctx, do)
	if err != nil {
		return err
	}

	// 由于网关服务没有与Spoved进行解耦，这里需要对部署Spoved的情况进行特殊处理
	// 在部署Spoved服务以后，需要调整ingress的入口
	proj := do.CompileOrder.Project
	if proj.Creator.Username == "root" && proj.ProjectName == "spoved-ingress" {
		_, err = k3s_service.UpdateSpovedIngress(ctx, do)
	}

	// 如果不是以小流量模式进行部署的，还需要下线掉旧版本的Service和Deployment
	if !do.Miniflow {
		err = k3s_service.ClearDeploymentByDeployOrder(ctx, do)
		if err != nil {
			return fmt.Errorf("clear deployment error: %w", err)
		}
		err = k3s_service.ClearServicesByDeployOrder(ctx, do)
		if err != nil {
			return fmt.Errorf("clear services error: %w", err)
		}
	}

	return err
}
