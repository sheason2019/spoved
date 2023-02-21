package compile

import (
	"github.com/gin-gonic/gin"
)

type CompileApi interface {
	PostCompileOrder(ctx *gin.Context, payload CompileOrder) CompileOrder
	GetCompileOrders(ctx *gin.Context, payload GetCompileOrdersPayload) GetCompileOrdersResponse
}
type _compileApiDefinition struct {
	GET_COMPILE_ORDERS_PATH string
	POST_COMPILE_ORDER_PATH string
}

var CompileApiDefinition = _compileApiDefinition{
	GET_COMPILE_ORDERS_PATH: "/CompileApi.CompileOrders",
	POST_COMPILE_ORDER_PATH: "/CompileApi.CompileOrder",
}
