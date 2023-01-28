package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/exceptions/exception"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	login_service "github.com/sheason2019/spoved/libs/service/account/login"
)

func UserMiddleware(ctx *gin.Context) {
	// 获取Token
	authorization := ctx.GetHeader("Authorization")
	// 不存在则不获取用户信息
	if len(authorization) == 0 {
		return
	}

	// 解析Token获取用户身份
	claims, e := login_service.ParseJwt(authorization)
	if e != nil {
		return
	}

	// 获取用户信息
	usr := &claims.User
	ctx.Set("user", usr)
}

// 从Token中获取当前请求的用户信息
func GetCurrentUser(ctx *gin.Context) (*ent.User, *exception.Exception) {
	value, exist := ctx.Get("user")
	if !exist {
		return nil, nil
	}

	usr, ok := value.(*ent.User)
	if !ok {
		return nil, exception.New(errors.New("获取用户信息失败"))
	}

	// 在数据库中索引用户
	usr, e := account_service.FindUserByUsername(usr.Username)
	if e != nil {
		return nil, e.Wrap()
	}
	if usr == nil {
		return nil, exception.New(errors.New("指定的用户信息不存在"))
	}

	return usr, nil
}

// 从Token中获取当前请求的用户信息，如果信息不存在则Panic
func MustGetCurrentUser(ctx *gin.Context) *ent.User {
	usr, e := GetCurrentUser(ctx)
	if e != nil {
		e.Panic()
	}
	if usr == nil {
		exception.New(errors.New("获取当前请求的身份信息失败")).Panic()
	}
	return usr
}
