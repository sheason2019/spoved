package regist_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateUser(ctx context.Context, username, password, salt string) (*dao.User, error) {
	client := dbc.GetClient()
	usr := &dao.User{
		Username:     username,
		PasswordHash: password,
		PasswordSalt: salt,
	}

	err := client.WithContext(ctx).Save(usr).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return usr, nil
}
