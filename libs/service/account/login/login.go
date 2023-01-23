package login_service

import (
	"errors"

	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	"github.com/sheason2019/spoved/libs/service/account/password"
)

func Login(info *account.AccountInfo) (string, *exception.Exception) {
	// 将用户上传的PWD解析成明文
	e := password.DecodePassword(info)
	if e != nil {
		return "", e.Wrap()
	}
	// 获取指定的用户信息
	usr, e := account_service.FindUserByUsername(info.Username)
	if e != nil {
		return "", e.Wrap()
	}
	if usr == nil {
		return "", exception.New(errors.New("用户名或密码错误"))
	}
	// 验证密码是否正确
	if password.StringHash(info.Password+usr.PasswordSalt) != usr.PasswordHash {
		return "", exception.New(errors.New("用户名或密码错误"))
	}

	// 登录成功后创建JWT
	token, e := GenerateJwt(usr)
	if e != nil {
		return "", e.Wrap()
	}

	return token, nil
}
