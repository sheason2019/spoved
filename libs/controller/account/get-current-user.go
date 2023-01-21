package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
)

func (AccountController) GetCurrentUser(ctx *gin.Context) account.User {
	return account.User{}
}
