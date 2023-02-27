package k3s_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 获取指定Project的服务状态
func FindProjectServices(ctx context.Context, proj *dao.Project) (*v1.ServiceList, error) {
	selector := meta_v1.LabelSelector{
		MatchLabels: map[string]string{
			"owner":       proj.Creator.Username,
			"projectName": proj.ProjectName,
		},
	}

	return clientSet.CoreV1().Services("default").List(ctx, meta_v1.ListOptions{
		LabelSelector: meta_v1.FormatLabelSelector(&selector),
	})
}
