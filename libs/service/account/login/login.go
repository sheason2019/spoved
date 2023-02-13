package login_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	"github.com/sheason2019/spoved/libs/service/account/password"
)

func Login(ctx context.Context, info *account.AccountInfo) (string, error) {
	// 将用户上传的PWD解析成明文
	err := password.DecodePassword(info)
	if err != nil {
		return "", err
	}
	// 获取指定的用户信息
	usr, err := account_service.FindUserByUsername(ctx, info.Username)
	if err != nil {
		return "", err
	}
	if usr == nil {
		return "", errors.WithStack(errors.New("用户名或密码错误"))
	}
	// 验证密码是否正确
	if password.StringHash(info.Password+usr.PasswordSalt) != usr.PasswordHash {
		return "", errors.WithStack(errors.New("用户名或密码错误"))
	}

	// 登录成功后创建JWT
	token, err := GenerateJwt(usr)
	if err != nil {
		return "", err
	}

	return token, nil
}
