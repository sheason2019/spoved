package deploy_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
	"github.com/sheason2019/spoved/libs/middleware"
	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	"github.com/sheason2019/spoved/libs/transfer"
)

func (deployController) GetDeployOrders(ctx *gin.Context, payload deploy.GetDeployOrdersPayload) deploy.GetDeployOrdersResponse {
	fmt.Println("payload", payload)
	if payload.ProjectId == 0 {
		panic("ProjectID不能为0")
	}

	orderDaos, count, err := deploy_service.FindDeployOrders(ctx, payload.ProjectId, payload.Page, payload.PageSize)
	if err != nil {
		panic(err)
	}

	pagination := common.Pagination{
		Page:       payload.Page,
		PageSize:   payload.PageSize,
		ItemCounts: count,
	}

	orders := make([]deploy.DeployOrder, len(orderDaos))
	for i, orderDao := range orderDaos {
		orders[i] = *transfer.DeployOrderToIdl(&orderDao)
	}

	return deploy.GetDeployOrdersResponse{
		Pagination: pagination,
		Records:    orders,
	}
}

func bindGetDeployOrders(r *gin.Engine) {
	r.GET(deploy.DeployApiDefinition.GET_DEPLOY_ORDERS_PATH, func(ctx *gin.Context) {
		payload := middleware.GetProps[deploy.GetDeployOrdersPayload](ctx)
		ctx.JSON(200, dc.GetDeployOrders(ctx, *payload))
	})
}
