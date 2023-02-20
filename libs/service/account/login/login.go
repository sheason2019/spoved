package login_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/initial"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	"github.com/sheason2019/spoved/libs/service/account/password"
)

func Login(ctx context.Context, info *account.AccountInfo) (string, error) {
	// 将用户上传的PWD解析成明文
	err := password.DecodePassword(info)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	usr, err := account_service.FindUserByUsername(ctx, info.Username)
	if err != nil {
		return "", err
	}

	if info.Username == "root" {
		err = rootLogin(ctx, usr, info)
	} else {
		err = userLogin(ctx, usr, info)
	}

	if err != nil {
		return "", err
	}

	return GenerateJwt(usr)
}

// Root用户登录
func rootLogin(ctx context.Context, usr *dao.User, info *account.AccountInfo) error {
	rootPwd := initial.GetRootPassword()
	if info.Password != rootPwd {
		return errors.WithStack(errors.New("用户名或密码错误"))
	}

	return nil
}

// 普通用户登录
func userLogin(ctx context.Context, usr *dao.User, info *account.AccountInfo) error {
	// 获取指定的用户信息
	if usr == nil {
		return errors.WithStack(errors.New("用户名或密码错误"))
	}
	// 验证密码是否正确
	if password.StringHash(info.Password+usr.PasswordSalt) != usr.PasswordHash {
		return errors.WithStack(errors.New("用户名或密码错误"))
	}

	return nil
}
