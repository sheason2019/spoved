package k3s_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/api/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 寻找Project下所有已创建的部署
func FindProjectDeployments(ctx context.Context, proj *dao.Project) (*v1.DeploymentList, error) {
	selector := meta_v1.LabelSelector{
		MatchLabels: map[string]string{
			"owner":       proj.Creator.Username,
			"projectName": proj.ProjectName,
		},
	}

	return clientSet.AppsV1().Deployments("default").List(ctx, meta_v1.ListOptions{
		LabelSelector: meta_v1.FormatLabelSelector(&selector),
	})
}
