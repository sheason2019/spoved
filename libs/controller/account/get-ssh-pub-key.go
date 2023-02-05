package account_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/middleware"
	ssh_key "github.com/sheason2019/spoved/libs/service/account/ssh-key"
)

func (AccountController) GetSshPubKey(ctx *gin.Context) account.GetSshPubKeyResponse {
	currentUser := middleware.MustGetCurrentUser(ctx)

	pair, err := ssh_key.GetSshKeyForce(currentUser)
	if err != nil {
		panic(err)
	}

	return account.GetSshPubKeyResponse{
		PubKey: pair.PublicKey,
	}
}
