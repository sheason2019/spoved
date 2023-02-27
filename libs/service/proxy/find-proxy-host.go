package proxy_service

import (
	"context"

	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func FindProxyHost(ctx context.Context, username, projectName string) (string, error) {
	proj, err := project_service.FindProject(ctx, "root", "spoved-fe")
	if err != nil {
		return "", err
	}

	do, err := deploy_service.FindLatestDeployOrder(ctx, proj)
	if err != nil {
		return "", err
	}

	return do.ServiceName, nil
}
