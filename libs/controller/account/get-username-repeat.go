package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
	"github.com/sheason2019/spoved/libs/utils"
)

func (AccountController) GetUsernameRepeat(ctx *gin.Context, payload account.GetUsernameRepeatPayload) account.GetUsernameRepeatResponse {
	repeat, e := regist_service.CheckUsernameRepeat(payload.Name)
	if e != nil {
		e.Panic()
	}
	return account.GetUsernameRepeatResponse{
		Repeat: repeat,
	}
}

func bindGetUsernameRepeat(r *gin.Engine) {
	r.GET(account.AccountServiceDefinition.USERNAME_REPEAT_PATH, func(ctx *gin.Context) {
		props := utils.GetProps[account.GetUsernameRepeatPayload](ctx)
		ctx.JSON(200, ac.GetUsernameRepeat(ctx, *props))
	})
}
