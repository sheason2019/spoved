package router

import (
	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
	compile_controller "github.com/sheason2019/spoved/libs/controller/compile"
	deploy_controller "github.com/sheason2019/spoved/libs/controller/deploy"
	project_controller "github.com/sheason2019/spoved/libs/controller/project"
	"github.com/sheason2019/spoved/libs/middleware"
)

// API Router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.UserMiddleware)

	account_controller.BindController(r)
	project_controller.BindController(r)
	compile_controller.BindController(r)
	deploy_controller.BindController(r)

	return r
}
