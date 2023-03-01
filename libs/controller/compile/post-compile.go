package compile_controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/middleware"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/transfer"
	"github.com/sheason2019/spoved/libs/utils"
)

// 创建编译工单
func (compileController) PostCompileOrder(ctx *gin.Context, payload compile.CompileOrder) compile.CompileOrder {
	currentUser := middleware.MustGetCurrentUser(ctx)

	proj, err := project_service.FindProjectById(ctx, payload.ProjectId)
	if err != nil {
		panic(err)
	}

	// 权限校验
	err = compile_service.ValidateOperator(ctx, proj, currentUser)
	if err != nil {
		panic(err)
	}

	// fetch 下一个版本的版本号
	nv, err := compile_service.FindNextVersionForProject(ctx, int(proj.ID), payload.Version)
	if err != nil {
		panic(err)
	}

	// 创建环境变量
	envMap := utils.StringToMap(payload.Env)
	if payload.Production {
		envMap["PRODUCTION"] = "true"
	} else {
		envMap["PRODUCTION"] = "false"
	}

	// TODO: 校验输入的环境变量是否合法

	// 创建编译工单
	order := &dao.CompileOrder{
		Image:      payload.Image,
		Version:    nv,
		Branch:     payload.Branch,
		Project:    *proj,
		Operator:   *currentUser,
		Production: payload.Production,
		Env:        envMap,
	}

	err = compile_service.CreateCompileOrder(ctx, order)
	if err != nil {
		panic(err)
	}

	go compile_service.CompileRun(context.TODO(), order)

	return *transfer.CompileOrderToIdl(order)
}

func bindPostCompile(r gin.IRoutes) {
	r.POST(compile.CompileApiDefinition.POST_COMPILE_ORDER_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[compile.CompileOrder](ctx)
		ctx.JSON(200, cc.PostCompileOrder(ctx, *props))
	})
}
