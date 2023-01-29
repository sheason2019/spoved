package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
	"github.com/sheason2019/spoved/libs/utils"
)

// 获取登录加密信息
func (AccountController) GetAccountCrypto(c *gin.Context) account.AccountCrypto {
	// 加密盐是 32 位随机字符串
	salt := utils.RandomStr(32)
	k := crypto_service.MustGetRsaPair()
	return account.AccountCrypto{
		Salt:   salt,
		PubKey: k.PubKey,
	}
}

func bindGetAccountCrypto(r *gin.Engine) {
	r.GET(account.AccountApiDefinition.ACCOUNT_CRYPTO_PATH, func(ctx *gin.Context) {
		ctx.JSON(200, ac.GetAccountCrypto(ctx))
	})
}
