package deploy_controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
	"github.com/sheason2019/spoved/libs/middleware"
)

func (deployController) PostDeployOrder(ctx *gin.Context, payload deploy.DeployOrder) deploy.DeployOrder {
	fmt.Printf("payload: %+v\n", payload)

	return deploy.DeployOrder{}
}

func bindPostDeployOrder(r gin.IRoutes) {
	r.POST(deploy.DeployApiDefinition.POST_DEPLOY_ORDER_PATH, func(ctx *gin.Context) {
		paylaod := middleware.GetProps[deploy.DeployOrder](ctx)
		ctx.JSON(200, dc.PostDeployOrder(ctx, *paylaod))
	})
}
