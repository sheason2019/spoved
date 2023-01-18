package main

import (
	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
)

func main() {
	r := gin.Default()

	account_controller.BindController(r)

	r.Run()
}
