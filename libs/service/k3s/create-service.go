package k3s_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateServiceByProject(ctx context.Context, proj *dao.Project) error {
	// 根据Project的信息对Service名称进行初始化
	serviceName := "service-proj-id-" + fmt.Sprint(proj.ID)
	// 创建Selector
	selector := meta_v1.LabelSelector{
		MatchLabels: map[string]string{
			"owner":       proj.Creator.Username,
			"projectName": proj.ProjectName,
		},
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
		proj.ServiceName = svc.Name
	} else {
		// 否则创建一个Service，并将ServiceName赋给Project
		svc := proj.GenerateService(serviceName)
		svc, err := postService(ctx, svc)
		if err != nil {
			return err
		}
		proj.ServiceName = svc.Name
	}

	// 保存Project
	return dbc.DB.WithContext(ctx).Save(proj).Error
}

func postService(ctx context.Context, svc *v1.Service) (*v1.Service, error) {
	return clientSet.CoreV1().Services("default").Create(ctx, svc, meta_v1.CreateOptions{})
}
