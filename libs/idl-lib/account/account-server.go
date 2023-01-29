package account

import (
	"github.com/gin-gonic/gin"
)

type AccountApi interface {
	GetAccountCrypto(ctx *gin.Context) AccountCrypto
	Login(ctx *gin.Context, account AccountInfo) LoginResponse
	Regist(ctx *gin.Context, account AccountInfo)
	GetUsernameRepeat(ctx *gin.Context, payload GetUsernameRepeatPayload) GetUsernameRepeatResponse
	GetCurrentUser(ctx *gin.Context) User
}
type _accountApiDefinition struct {
	ACCOUNT_CRYPTO_PATH  string
	LOGIN_PATH           string
	REGIST_PATH          string
	USERNAME_REPEAT_PATH string
	CURRENT_USER_PATH    string
}

var AccountApiDefinition = _accountApiDefinition{
	ACCOUNT_CRYPTO_PATH:  "/AccountApi.AccountCrypto",
	LOGIN_PATH:           "/AccountApi.Login",
	REGIST_PATH:          "/AccountApi.Regist",
	USERNAME_REPEAT_PATH: "/AccountApi.UsernameRepeat",
	CURRENT_USER_PATH:    "/AccountApi.CurrentUser",
}
