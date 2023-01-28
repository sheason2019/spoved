package project_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

type ProjectController struct{}

var pc project.ProjectController = ProjectController{}

func BindProjectController(r *gin.Engine) {
	bindPostProject(r)
}
