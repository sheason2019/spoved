package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
)

type compileController struct{}

var cc compile.CompileApi = compileController{}

func BindController(r *gin.Engine) {
	bindPostCompile(r)
}
