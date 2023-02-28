package transfer

import (
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
)

func DeployOrderToIdl(do *dao.DeployOrder) *deploy.DeployOrder {
	return &deploy.DeployOrder{
		Id:           int(do.ID),
		Image:        do.Image,
		CreateAt:     int(do.CreatedAt.Unix()),
		Operator:     do.Operator.Username,
		StatusCode:   do.StatusCode,
		CompileOrder: *CompileOrderToIdl(&do.CompileOrder),
		Miniflow:     do.Miniflow,
	}
}
