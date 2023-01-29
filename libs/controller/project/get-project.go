package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

func (ProjectController) GetProject(ctx *gin.Context, payload project.GetProjectPayload) project.Project {

	return project.Project{}
}
