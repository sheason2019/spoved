package proxy_service

import (
	"context"
	"fmt"

	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func findProxyHostInfo(ctx context.Context, username, projectName string) (*DebounceHostInfo, error) {
	proj, err := project_service.FindProject(ctx, "root", "spoved-fe")
	if err != nil {
		return nil, err
	}

	orders, err := deploy_service.FindRunningDeployOrders(ctx, proj)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("error: DeployOrder is nil on FindProxyHost")
	}

	info := &DebounceHostInfo{}

	miniflowMatches := []HostMatch{}
	info.Miniflow = miniflowMatches

	for _, order := range orders {
		match := HostMatch{
			HostPath:    order.ServiceName,
			HeaderMatch: order.HeaderPair,
		}
		if order.Miniflow {
			miniflowMatches = append(miniflowMatches, match)
		} else {
			info.Online = match
		}
	}

	return info, nil
}
