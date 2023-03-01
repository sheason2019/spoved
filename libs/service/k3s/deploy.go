package k3s_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 启动服务
func CreateDeploymentByDeployOrder(ctx context.Context, do *dao.DeployOrder) error {
	deployment := do.GenerateDeployment("deployment-order-id-" + fmt.Sprint(do.ID))

	_, err := clientSet.AppsV1().Deployments("default").Create(ctx, deployment, v1.CreateOptions{})

	return err
}

func ClearDeploymentByDeployOrder(ctx context.Context, do *dao.DeployOrder) error {
	// 获取项目下的所有Deployment
	deploys, err := FindProjectDeployments(ctx, &do.CompileOrder.Project)
	if err != nil {
		return fmt.Errorf("clear deploy: find project deploys error: %w", err)
	}

	// 获取应当保留的版本
	currentV := do.CompileOrder.Version

	// 对所有应当下线的Deployment执行删除操作（通常被下线的Deployment的个数为1）
	for _, deploy := range deploys.Items {
		if deploy.Labels["version"] != currentV && deploy.Labels["miniflow"] != "true" {
			err = clientSet.AppsV1().Deployments("default").Delete(ctx, deploy.Name, v1.DeleteOptions{})
			if err != nil {
				return fmt.Errorf("error when delete deploy:%w", err)
			}
		}
	}

	return nil
}
