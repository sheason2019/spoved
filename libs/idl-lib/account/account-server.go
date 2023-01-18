package account

import (
	"github.com/gin-gonic/gin"
)

type AccountService interface {
	GetAccountCrypto(ctx *gin.Context) AccountCrypto
	Login(ctx *gin.Context, account AccountInfo) LoginResponse
	Regist(ctx *gin.Context, account AccountInfo)
}
type _accountServiceDefinition struct {
	REGIST_PATH         string
	ACCOUNT_CRYPTO_PATH string
	LOGIN_PATH          string
}

var AccountServiceDefinition = _accountServiceDefinition{
	REGIST_PATH:         "/AccountService.Regist",
	ACCOUNT_CRYPTO_PATH: "/AccountService.AccountCrypto",
	LOGIN_PATH:          "/AccountService.Login",
}
