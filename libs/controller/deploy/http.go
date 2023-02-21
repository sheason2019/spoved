package deploy_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
)

type deployController struct{}

var dc deploy.DeployApi = deployController{}

func BindController(r *gin.Engine) {
	bindGetDeployOrders(r)
}
