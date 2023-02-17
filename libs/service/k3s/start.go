package k3s_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 启动服务
func Start(ctx context.Context, do *dao.DeployOrder) error {
	deployment := do.GenerateDeployment("deployment-order-id-" + fmt.Sprint(do.ID))

	_, err := clientSet.AppsV1().Deployments("default").Create(ctx, deployment, v1.CreateOptions{})

	return err
}
