package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
)

func (compileController) GetCompileRecords(ctx *gin.Context, pagination common.Pagination) compile.GetCompileRecordResponse {
	return compile.GetCompileRecordResponse{}
}
