package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/middleware"
	login_service "github.com/sheason2019/spoved/libs/service/account/login"
)

func (AccountController) Login(c *gin.Context, info account.AccountInfo) account.LoginResponse {
	token, e := login_service.Login(&info)
	if e != nil {
		panic(e)
	}

	return account.LoginResponse{
		Token: token,
	}
}

func bindLogin(r *gin.Engine) {
	r.POST(account.AccountApiDefinition.LOGIN_PATH, func(ctx *gin.Context) {
		props := middleware.GetProps[account.AccountInfo](ctx)
		ctx.JSON(200, ac.Login(ctx, *props))
	})
}
