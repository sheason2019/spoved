package project

import (
	"github.com/gin-gonic/gin"
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type ProjectApi interface {
	GetProjects(ctx *gin.Context, pagination common.Pagination) GetProjectsResponse
	PostProject(ctx *gin.Context, project Project) Project
}
type _projectApiDefinition struct {
	PROJECTS_PATH string
	PROJECT_PATH  string
}

var ProjectApiDefinition = _projectApiDefinition{
	PROJECTS_PATH: "/ProjectApi.Projects",
	PROJECT_PATH:  "/ProjectApi.Project",
}
