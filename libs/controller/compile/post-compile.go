package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/middleware"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/transfer"
)

// 创建编译工单
func (compileController) PostCompile(ctx *gin.Context, payload compile.CompileRecord) compile.CompileRecord {
	currentUser := middleware.MustGetCurrentUser(ctx)

	proj, err := project_service.FindProjectById(payload.ProjectId)
	if err != nil {
		panic(err)
	}

	nv, err := compile_service.FindNextVersionForProject(proj.ID, payload.Version)
	if err != nil {
		panic(err)
	}

	record, err := compile_service.Compile(ctx, payload.Image, nv, payload.Branch, proj, currentUser)
	if err != nil {
		panic(err)
	}

	return *transfer.CompileRecordToIdl(ctx, record)
}

func bindPostCompile(r *gin.Engine) {
	r.POST(compile.CompileApiDefinition.POST_COMPILE_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[compile.CompileRecord](ctx)
		ctx.JSON(200, cc.PostCompile(ctx, *props))
	})
}
