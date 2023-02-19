package deploy_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
)

type DeployContext struct {
	context.Context
	DeployOrder *dao.DeployOrder
}
