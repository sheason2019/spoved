package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/middleware"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/transfer"
)

// 创建编译工单
func (compileController) PostCompileOrder(ctx *gin.Context, payload compile.CompileOrder) compile.CompileOrder {
	currentUser := middleware.MustGetCurrentUser(ctx)

	proj, err := project_service.FindProjectById(ctx, payload.ProjectId)
	if err != nil {
		panic(err)
	}

	nv, err := compile_service.FindNextVersionForProject(ctx, int(proj.ID), payload.Version)
	if err != nil {
		panic(err)
	}

	order := &dao.CompileOrder{
		Image:    payload.Image,
		Version:  nv,
		Branch:   payload.Branch,
		Project:  *proj,
		Operator: *currentUser,
	}
	err = compile_service.CreateCompileOrder(ctx, order)
	if err != nil {
		panic(err)
	}

	return *transfer.CompileOrderToIdl(order)
}

func bindPostCompile(r gin.IRoutes) {
	r.POST(compile.CompileApiDefinition.POST_COMPILE_ORDER_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[compile.CompileOrder](ctx)
		ctx.JSON(200, cc.PostCompileOrder(ctx, *props))
	})
}
