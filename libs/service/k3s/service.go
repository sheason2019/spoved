package k3s_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateServiceByDeployOrder(ctx context.Context, do *dao.DeployOrder) error {
	// 根据Project的信息对Service名称进行初始化
	serviceName := "service-deploy-order-id-" + fmt.Sprint(do.ID)
	// 创建Selector
	selector := meta_v1.LabelSelector{
		MatchLabels: do.GenerateSelector(),
	}

	// 寻找指定的服务，判断是否已经启动
	services, err := clientSet.CoreV1().Services("default").List(ctx, meta_v1.ListOptions{
		LabelSelector: meta_v1.FormatLabelSelector(&selector),
	})
	if err != nil {
		return err
	}

	if len(services.Items) > 0 {
		// 如果Project对应的Service已存在，则将ServiceName赋给Project
		svc := services.Items[0]
		do.ServiceName = svc.Name
	} else {
		// 否则创建一个Service，并将ServiceName赋给Project
		svc := do.GenerateService(serviceName)
		svc, err := clientSet.CoreV1().Services("default").Create(ctx, svc, meta_v1.CreateOptions{})
		if err != nil {
			return err
		}
		do.ServiceName = svc.Name
	}

	// 保存Project
	return dbc.DB.WithContext(ctx).Save(do).Error
}

func ClearServicesByDeployOrder(ctx context.Context, do *dao.DeployOrder) error {
	// 获取项目下的所有Service
	services, err := FindProjectServices(ctx, &do.CompileOrder.Project)
	if err != nil {
		return fmt.Errorf("error find porjectServices: %w", err)
	}

	// 获取应当保留的版本
	currentV := do.CompileOrder.Version

	// 删除非指定版本的Service
	for _, service := range services.Items {
		if service.Labels["version"] != currentV {
			err = clientSet.CoreV1().Services("default").Delete(ctx, service.Name, meta_v1.DeleteOptions{})
			if err != nil {
				return fmt.Errorf("error occur when delete service: %w", err)
			}
		}
	}

	return nil
}
