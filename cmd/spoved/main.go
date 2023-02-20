package main

import (
	"github.com/sheason2019/spoved/libs/env"
	"github.com/sheason2019/spoved/libs/initial"
	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	// 在stdout中展示root密码
	initial.ShowPassword()

	r := router.SetupRouter()

	var port string
	if env.IS_PRODUCT {
		port = ":80"
	} else {
		port = ":8080"
	}

	r.Run(port)
}
