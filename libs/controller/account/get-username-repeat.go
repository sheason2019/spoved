package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/middleware"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
)

func (AccountController) GetUsernameRepeat(ctx *gin.Context, payload account.GetUsernameRepeatPayload) account.GetUsernameRepeatResponse {
	usr, e := account_service.FindUserByUsername(payload.Name)
	if e != nil {
		panic(e)
	}
	return account.GetUsernameRepeatResponse{
		Repeat: usr != nil,
	}
}

func bindGetUsernameRepeat(r gin.IRoutes) {
	r.GET(account.AccountApiDefinition.GET_USERNAME_REPEAT_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[account.GetUsernameRepeatPayload](ctx)
		ctx.JSON(200, ac.GetUsernameRepeat(ctx, *props))
	})
}
