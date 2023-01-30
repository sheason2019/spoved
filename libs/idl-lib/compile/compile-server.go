package compile

import (
	"github.com/gin-gonic/gin"
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type CompileApi interface {
	PostCompile(ctx *gin.Context, payload CompileRecord) CompileRecord
	GetCompileRecords(ctx *gin.Context, pagination common.Pagination) GetCompileRecordResponse
}
type _compileApiDefinition struct {
	POST_COMPILE_PATH        string
	GET_COMPILE_RECORDS_PATH string
}

var CompileApiDefinition = _compileApiDefinition{
	POST_COMPILE_PATH:        "/CompileApi.Compile",
	GET_COMPILE_RECORDS_PATH: "/CompileApi.CompileRecords",
}
