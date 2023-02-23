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
	GetSshPubKey(ctx *gin.Context) GetSshPubKeyResponse
}
type _accountApiDefinition struct {
	GET_SSH_PUB_KEY_PATH     string
	GET_ACCOUNT_CRYPTO_PATH  string
	LOGIN_PATH               string
	REGIST_PATH              string
	GET_USERNAME_REPEAT_PATH string
	GET_CURRENT_USER_PATH    string
}

var AccountApiDefinition = _accountApiDefinition{
	GET_SSH_PUB_KEY_PATH:     "/AccountApi.SshPubKey",
	GET_ACCOUNT_CRYPTO_PATH:  "/AccountApi.AccountCrypto",
	LOGIN_PATH:               "/AccountApi.Login",
	REGIST_PATH:              "/AccountApi.Regist",
	GET_USERNAME_REPEAT_PATH: "/AccountApi.UsernameRepeat",
	GET_CURRENT_USER_PATH:    "/AccountApi.CurrentUser",
}
