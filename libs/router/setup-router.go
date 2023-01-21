package router

import (
	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
	"github.com/sheason2019/spoved/libs/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Recovery)
	r.Use(middleware.DataLog)

	account_controller.BindController(r)

	return r
}
