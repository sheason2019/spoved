package account_controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	login_service "github.com/sheason2019/spoved/libs/service/account/login"
)

func (AccountController) GetCurrentUser(ctx *gin.Context) account.User {
	// 获取Token
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 {
		exception.New(errors.New("请求中没有包含必须的用户信息")).Panic()
	}

	// 解析Token获取用户身份
	claims, e := login_service.ParseJwt(authorization)
	if e != nil {
		e.Panic()
	}

	// 获取用户信息
	usr := &claims.User
	// 在数据库中索引用户
	usr, e = account_service.FindUserByUsername(usr.Username)
	if e != nil {
		e.Panic()
	}
	if usr == nil {
		exception.New(errors.New("指定的用户信息不存在")).Panic()
	}

	return account.User{
		Username: usr.Username,
	}
}

func bindGetCurrentUser(r *gin.Engine) {
	r.GET(account.AccountServiceDefinition.CURRENT_USER_PATH, func(ctx *gin.Context) {
		ctx.JSON(200, ac.GetCurrentUser(ctx))
	})
}
