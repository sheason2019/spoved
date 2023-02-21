package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/middleware"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	"github.com/sheason2019/spoved/libs/transfer"
)

func (compileController) GetCompileOrders(ctx *gin.Context, payload compile.GetCompileOrdersPayload) compile.GetCompileOrdersResponse {
	recordDaos, count, err := compile_service.FindCompileOrders(ctx, payload.ProjectId, payload.Page, payload.PageSize)
	if err != nil {
		panic(err)
	}

	pagination := common.Pagination{
		Page:       payload.Page,
		PageSize:   payload.PageSize,
		ItemCounts: count,
	}

	records := make([]compile.CompileOrder, len(recordDaos))

	for i, recordDao := range recordDaos {
		records[i] = *transfer.CompileOrderToIdl(&recordDao)
	}

	return compile.GetCompileOrdersResponse{
		Records:    records,
		Pagination: pagination,
	}
}

func bindGetCompileOrders(r gin.IRoutes) {
	r.GET(compile.CompileApiDefinition.GET_COMPILE_ORDERS_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[compile.GetCompileOrdersPayload](ctx)
		ctx.JSON(200, cc.GetCompileOrders(ctx, *props))
	})
}
