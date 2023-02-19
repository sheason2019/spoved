package regist_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	password_service "github.com/sheason2019/spoved/libs/service/account/password"
)

func Regist(ctx context.Context, username, password string) (*dao.User, error) {
	// 将用户的密文密码转换为Salt+Hash的组合
	cipherPwd, salt := password_service.EncodePassword(password)
	// 创建用户
	usr, e := CreateUser(ctx, username, cipherPwd, salt)
	if e != nil {
		return nil, e
	}

	return usr, nil
}
