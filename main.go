package main

import (
	"context"

	"github.com/sheason2019/spoved/libs/router"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

func main() {
	// 初始化前端服务
	k3s_service.InitSpovedFe(context.TODO())

	r := router.SetupRouter()

	r.Run(":80")
}
