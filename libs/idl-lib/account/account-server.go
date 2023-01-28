package account

import (
	"github.com/gin-gonic/gin"
)

type AccountService interface {
	GetAccountCrypto(ctx *gin.Context) AccountCrypto
	Login(ctx *gin.Context, account AccountInfo) LoginResponse
	Regist(ctx *gin.Context, account AccountInfo)
	GetUsernameRepeat(ctx *gin.Context, payload GetUsernameRepeatPayload) GetUsernameRepeatResponse
	GetCurrentUser(ctx *gin.Context) User
}
type _accountServiceDefinition struct {
	ACCOUNT_CRYPTO_PATH  string
	LOGIN_PATH           string
	REGIST_PATH          string
	USERNAME_REPEAT_PATH string
	CURRENT_USER_PATH    string
}

var AccountServiceDefinition = _accountServiceDefinition{
	ACCOUNT_CRYPTO_PATH:  "/AccountService.AccountCrypto",
	LOGIN_PATH:           "/AccountService.Login",
	REGIST_PATH:          "/AccountService.Regist",
	USERNAME_REPEAT_PATH: "/AccountService.UsernameRepeat",
	CURRENT_USER_PATH:    "/AccountService.CurrentUser",
}
