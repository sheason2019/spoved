package main

import (
	"github.com/sheason2019/spoved/libs/initial"
	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	// 在stdout中展示root密码
	initial.ShowPassword()

	r := router.SetupRouter()

	r.Run(":8080")
}
