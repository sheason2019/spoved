package project_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	"github.com/sheason2019/spoved/libs/middleware"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func (ProjectController) PostProject(ctx *gin.Context, proj project.Project) project.Project {
	usr := middleware.MustGetCurrentUser(ctx)

	entProj, e := project_service.CreateProject(&proj, usr)
	if e != nil {
		panic(fmt.Sprintf("%+v", e))
	}

	return project.Project{
		Id:          entProj.ID,
		ProjectName: entProj.ProjectName,
		Describe:    entProj.Describe,
		GitUrl:      entProj.GitURL,
		Owner:       usr.Username,
	}
}

func bindPostProject(r *gin.Engine) {
	r.POST(project.ProjectApiDefinition.POST_PROJECT_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[project.Project](ctx)
		ctx.JSON(200, pc.PostProject(ctx, *props))
	})
}
