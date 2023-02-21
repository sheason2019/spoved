package deploy

import (
	"github.com/gin-gonic/gin"
)

type DeployApi interface {
	GetDeployOrders(ctx *gin.Context, payload GetDeployOrdersPayload) GetDeployOrdersResponse
}
type _deployApiDefinition struct {
	GET_DEPLOY_ORDERS_PATH string
}

var DeployApiDefinition = _deployApiDefinition{
	GET_DEPLOY_ORDERS_PATH: "/DeployApi.DeployOrders",
}
