package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

type ProjectController struct{}

var pc project.ProjectApi = ProjectController{}

func BindProjectController(r *gin.Engine) {
	bindPostProject(r)
	bindGetProjects(r)
	bindGetProject(r)
}
