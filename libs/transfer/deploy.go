package transfer

import (
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
)

func DeployOrderToIdl(do *dao.DeployOrder) *deploy.DeployOrder {
	return &deploy.DeployOrder{
		Id:             int(do.ID),
		ProjectId:      do.CompileOrder.ProjectID,
		Image:          do.Image,
		CreateAt:       int(do.CreatedAt.Unix()),
		Operator:       do.Operator.Username,
		CompileOrderId: do.CompileOrderID,
		StatusCode:     do.StatusCode,
	}
}
