package router

import (
	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	account_controller.BindController(r)

	return r
}
