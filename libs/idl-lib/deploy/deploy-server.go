package deploy

import (
	"github.com/gin-gonic/gin"
)

type DeployApi interface {
	GetCompileOrders(ctx *gin.Context, payload GetCompileOrdersPayload) GetCompileOrdersResponse
}
type _deployApiDefinition struct {
	GET_COMPILE_ORDERS_PATH string
}

var DeployApiDefinition = _deployApiDefinition{
	GET_COMPILE_ORDERS_PATH: "/DeployApi.CompileOrders",
}
