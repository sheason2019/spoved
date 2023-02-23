package compile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	"github.com/sheason2019/spoved/libs/service/images"
)

func (compileController) GetOptionalImages(ctx *gin.Context) compile.GetOptionalImagesResponse {
	return compile.GetOptionalImagesResponse{
		Images: images.OptionalCompileImages,
	}
}

func bindGetOptionalImages(r gin.IRoutes) {
	r.GET(compile.CompileApiDefinition.GET_OPTIONAL_IMAGES_PATH, func(ctx *gin.Context) {
		ctx.JSON(200, cc.GetOptionalImages(ctx))
	})
}
