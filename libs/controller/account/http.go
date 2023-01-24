package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
)

type AccountController struct{}

var ac account.AccountService = AccountController{}

func BindController(r *gin.Engine) {
	bindGetAccountCrypto(r)
	bindRegist(r)
	bindGetUsernameRepeat(r)
	bindLogin(r)
	bindGetCurrentUser(r)
}
