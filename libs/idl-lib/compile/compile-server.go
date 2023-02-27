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
	POST_COMPILE_ORDER_PATH  string
	GET_COMPILE_ORDERS_PATH  string
	GET_OPTIONAL_IMAGES_PATH string
}

var CompileApiDefinition = _compileApiDefinition{
	POST_COMPILE_ORDER_PATH:  "/CompileApi.CompileOrder",
	GET_COMPILE_ORDERS_PATH:  "/CompileApi.CompileOrders",
	GET_OPTIONAL_IMAGES_PATH: "/CompileApi.OptionalImages",
}
