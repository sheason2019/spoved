package project_controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	"github.com/sheason2019/spoved/libs/middleware"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/transfer"
)

func (ProjectController) GetProject(ctx *gin.Context, payload project.GetProjectPayload) project.Project {
	proj, e := project_service.FindProject(payload.Username, payload.ProjectName)
	if e != nil {
		e.Panic()
	}
	if proj == nil {
		exception.New(errors.New("指定的Project不存在")).Panic()
	}

	return transfer.ProjectToIdl(proj)
}

func bindGetProject(r *gin.Engine) {
	r.GET(project.ProjectApiDefinition.GET_PROJECT_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[project.GetProjectPayload](ctx)
		ctx.JSON(200, pc.GetProject(ctx, *props))
	})
}
