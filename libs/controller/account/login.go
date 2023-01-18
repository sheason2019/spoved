package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
)

func (AccountController) Login(c *gin.Context, info account.AccountInfo) account.LoginResponse {
	return account.LoginResponse{}
}
