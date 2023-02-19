package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	"github.com/sheason2019/spoved/libs/middleware"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/transfer"
	"github.com/sheason2019/spoved/libs/utils"
)

// 获取当前用户的Project信息
func (ProjectController) GetProjects(ctx *gin.Context, pagination common.Pagination) project.GetProjectsResponse {
	currentUser := middleware.MustGetCurrentUser(ctx)

	projs, e := project_service.FindProjectsByUser(ctx, currentUser, &pagination)
	if e != nil {
		panic(e)
	}
	count, e := project_service.CountProjectsByUser(ctx, currentUser)
	if e != nil {
		panic(e)
	}

	pagination.ItemCounts = int(count)

	projs_idl := utils.MapTo(projs, func(item dao.Project, index int) project.Project {
		return *transfer.ProjectToIdl(&item)
	})

	return project.GetProjectsResponse{
		Pagination: pagination,
		Projects:   projs_idl,
	}
}

func bindGetProjects(r gin.IRoutes) {
	r.GET(project.ProjectApiDefinition.GET_PROJECTS_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[common.Pagination](ctx)
		ctx.JSON(200, pc.GetProjects(ctx, *props))
	})
}
