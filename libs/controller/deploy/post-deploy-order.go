package deploy_controller

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/deploy"
	"github.com/sheason2019/spoved/libs/middleware"
	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	"github.com/sheason2019/spoved/libs/service/images"
	"github.com/sheason2019/spoved/libs/transfer"
)

func (deployController) PostDeployOrder(ctx *gin.Context, payload deploy.DeployOrder) deploy.DeployOrder {
	fmt.Printf("post deploy order payload: %+v\n", payload)
	currentUser := middleware.MustGetCurrentUser(ctx)

	// 校验当前用户是否有执行部署的权限
	err := deploy_service.ValidateDeployLimit(ctx, currentUser, &payload)
	if err != nil {
		panic(err)
	}
	// 校验镜像是否有效
	support := images.ValidateImage(payload.Image, "deploy")
	if !support {
		panic(fmt.Errorf("指定了无效的镜像：%s", payload.Image))
	}

	// 小流量头，仅进行小流量部署时写入该数据
	headerPair := map[string]string{}
	if payload.Miniflow {
		for _, pair := range payload.HeaderPairs {
			headerPair[pair.Header] = pair.Value
		}
	}

	// TODO: 校验输入的小流量头是否合法

	// 创建部署工单，首先需要根据相关的信息创建工单实体
	do := &dao.DeployOrder{
		CompileOrderID: payload.CompileOrder.Id,
		Image:          payload.Image,
		Operator:       *currentUser,
		Miniflow:       payload.Miniflow,
		HeaderPair:     headerPair,
	}
	err = deploy_service.CreateDeployOrder(ctx, do)
	if err != nil {
		panic(fmt.Errorf("创建DeployOrder失败: %w", err))
	}

	// 创建协程执行工单
	go deploy_service.DeployRun(context.TODO(), do)

	// 返回已创建的工单
	return *transfer.DeployOrderToIdl(do)
}

func bindPostDeployOrder(r gin.IRoutes) {
	r.POST(deploy.DeployApiDefinition.POST_DEPLOY_ORDER_PATH, func(ctx *gin.Context) {
		paylaod := middleware.GetProps[deploy.DeployOrder](ctx)
		ctx.JSON(200, dc.PostDeployOrder(ctx, *paylaod))
	})
}
