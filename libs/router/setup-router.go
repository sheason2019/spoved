package router

import (
	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
	"github.com/sheason2019/spoved/libs/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.Recovery)
	r.Use(middleware.DataLog)
	r.Use(middleware.UserMiddleware)

	account_controller.BindController(r)

	return r
}
