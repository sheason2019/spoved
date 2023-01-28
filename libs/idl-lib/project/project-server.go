package project

import (
	"github.com/gin-gonic/gin"
	common "github.com/sheason2019/spoved/libs/idl-lib/common"
)

type ProjectController interface {
	GetProjects(ctx *gin.Context, pagination common.Pagination) GetProjectsResponse
	PostProject(ctx *gin.Context, project Project) Project
}
type _projectControllerDefinition struct {
	PROJECT_PATH  string
	PROJECTS_PATH string
}

var ProjectControllerDefinition = _projectControllerDefinition{
	PROJECT_PATH:  "/ProjectController.Project",
	PROJECTS_PATH: "/ProjectController.Projects",
}
