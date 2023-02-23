package deploy_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
	"github.com/sheason2019/spoved/libs/service/images"
)

func (deployController) GetOptionalImages(ctx *gin.Context) deploy.GetOptionalImagesResponse {
	return deploy.GetOptionalImagesResponse{
		Images: images.OptionalDeployImages,
	}
}

func bindGetOptionalImages(r gin.IRoutes) {
	r.GET(deploy.DeployApiDefinition.GET_OPTIONAL_IMAGES_PATH, func(ctx *gin.Context) {
		ctx.JSON(200, dc.GetOptionalImages(ctx))
	})
}
