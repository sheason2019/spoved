package project

import (
	"github.com/gin-gonic/gin"
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type ProjectApi interface {
	GetProjects(ctx *gin.Context, pagination common.Pagination) GetProjectsResponse
	PostProject(ctx *gin.Context, project Project) Project
	GetProject(ctx *gin.Context, payload GetProjectPayload) Project
}
type _projectApiDefinition struct {
	GET_PROJECT_PATH  string
	GET_PROJECTS_PATH string
	POST_PROJECT_PATH string
}

var ProjectApiDefinition = _projectApiDefinition{
	GET_PROJECT_PATH:  "/ProjectApi.Project",
	GET_PROJECTS_PATH: "/ProjectApi.Projects",
	POST_PROJECT_PATH: "/ProjectApi.Project",
}
