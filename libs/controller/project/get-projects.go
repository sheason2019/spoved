package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

func (ProjectController) GetProjects(ctx *gin.Context, pagination common.Pagination) project.GetProjectsResponse {
	return project.GetProjectsResponse{}
}
