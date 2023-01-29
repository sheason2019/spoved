package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/middleware"
)

func (AccountController) GetCurrentUser(ctx *gin.Context) account.User {
	usr := middleware.MustGetCurrentUser(ctx)

	return account.User{
		Username: usr.Username,
	}
}

func bindGetCurrentUser(r *gin.Engine) {
	r.GET(account.AccountApiDefinition.CURRENT_USER_PATH, func(ctx *gin.Context) {
		ctx.JSON(200, ac.GetCurrentUser(ctx))
	})
}
