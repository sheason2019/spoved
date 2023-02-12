package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
)

type AccountController struct{}

var ac account.AccountApi = AccountController{}

func BindController(r gin.IRoutes) {
	bindGetAccountCrypto(r)
	bindRegist(r)
	bindGetUsernameRepeat(r)
	bindLogin(r)
	bindGetCurrentUser(r)
	bindGetSshPubKey(r)
}
