package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/middleware"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	"github.com/sheason2019/spoved/libs/transfer"
)

func (compileController) GetCompileRecords(ctx *gin.Context, payload compile.GetCompileRecordsPayload) compile.GetCompileRecordsResponse {
	recordDaos, count, err := compile_service.FindCompileRecords(ctx, payload.ProjectId, &payload.Pagination)
	if err != nil {
		panic(err)
	}

	pagination := payload.Pagination
	pagination.ItemCounts = count

	records := make([]compile.CompileRecord, len(recordDaos))

	for i, recordDao := range recordDaos {
		records[i] = *transfer.CompileRecordToIdl(&recordDao)
	}

	return compile.GetCompileRecordsResponse{
		Records:    records,
		Pagination: pagination,
	}
}

func bindGetCompileRecords(r gin.IRoutes) {
	r.GET(compile.CompileApiDefinition.GET_COMPILE_RECORDS_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[compile.GetCompileRecordsPayload](ctx)
		ctx.JSON(200, cc.GetCompileRecords(ctx, *props))
	})
}
