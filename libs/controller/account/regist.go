package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/middleware"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
)

func (AccountController) Regist(c *gin.Context, account account.AccountInfo) {
	// 注册逻辑 接受用户信息，校验并生成用户Record
	_, e := regist_service.Regist(c, &account)
	if e != nil {
		panic(e)
	}
}

func bindRegist(r gin.IRoutes) {
	r.POST(account.AccountApiDefinition.REGIST_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[account.AccountInfo](ctx)
		ac.Regist(ctx, *props)
		ctx.String(200, "OK")
	})
}
