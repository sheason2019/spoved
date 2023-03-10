package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	"github.com/sheason2019/spoved/libs/middleware"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func (ProjectController) PostProject(ctx *gin.Context, proj project.Project) project.Project {
	usr := middleware.MustGetCurrentUser(ctx)

	projDao, e := project_service.CreateProject(ctx, &proj, usr)
	if e != nil {
		panic(e)
	}

	return project.Project{
		Id:          int(projDao.ID),
		ProjectName: projDao.ProjectName,
		Describe:    projDao.Describe,
		GitUrl:      projDao.GitUrl,
		Owner:       usr.Username,
	}
}

func bindPostProject(r gin.IRoutes) {
	r.POST(project.ProjectApiDefinition.POST_PROJECT_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[project.Project](ctx)
		ctx.JSON(200, pc.PostProject(ctx, *props))
	})
}
