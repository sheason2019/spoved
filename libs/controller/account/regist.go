package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
	"github.com/sheason2019/spoved/libs/utils"
)

func (AccountController) Regist(c *gin.Context, account account.AccountInfo) {
	// 注册逻辑 接受用户信息，校验并生成用户Record
	_, e := regist_service.Regist(&account)
	if e != nil {
		e.Panic()
	}
}

func bindRegist(r *gin.Engine) {
	r.POST(account.AccountServiceDefinition.REGIST_PATH, func(ctx *gin.Context) {
		props := utils.GetProps[account.AccountInfo](ctx)
		ac.Regist(ctx, *props)
		ctx.String(200, "OK")
	})
}