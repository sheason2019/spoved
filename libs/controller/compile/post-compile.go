package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
)

func (compileController) PostCompile(ctx *gin.Context, payload compile.CompileRecord) compile.CompileRecord {
	return compile.CompileRecord{}
}
