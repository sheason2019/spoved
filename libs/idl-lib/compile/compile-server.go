package compile

import (
	"github.com/gin-gonic/gin"
)

type CompileApi interface {
	PostCompileOrder(ctx *gin.Context, payload CompileOrder) CompileOrder
	GetCompileOrders(ctx *gin.Context, payload GetCompileOrdersPayload) GetCompileOrdersResponse
	GetOptionalImages(ctx *gin.Context) GetOptionalImagesResponse
}
type _compileApiDefinition struct {
	GET_OPTIONAL_IMAGES_PATH string
	POST_COMPILE_ORDER_PATH  string
	GET_COMPILE_ORDERS_PATH  string
}

var CompileApiDefinition = _compileApiDefinition{
	GET_OPTIONAL_IMAGES_PATH: "/CompileApi.OptionalImages",
	POST_COMPILE_ORDER_PATH:  "/CompileApi.CompileOrder",
	GET_COMPILE_ORDERS_PATH:  "/CompileApi.CompileOrders",
}
