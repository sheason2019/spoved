package proxy_service

import (
	"context"
	"fmt"

	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func findProxyHost(ctx context.Context, username, projectName string) (string, error) {
	proj, err := project_service.FindProject(ctx, "root", "spoved-fe")
	if err != nil {
		return "", err
	}

	do, err := deploy_service.FindLatestDeployOrder(ctx, proj)
	if err != nil {
		return "", err
	}
	if do == nil {
		return "", fmt.Errorf("error: DeployOrder is nil on FindProxyHost")
	}

	return do.ServiceName, nil
}
