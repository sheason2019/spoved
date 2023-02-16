package deploy_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 部署逻辑执行
func DeployRun(ctx context.Context, do *dao.DeployOrder) {
	// 创建Deployment
	err := k3s_service.Start(ctx, do)
	if err != nil {
		fmt.Println("error::", err)
	}

	// 获取record信息
}
