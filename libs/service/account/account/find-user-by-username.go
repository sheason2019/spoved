package account_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindUserByUsername(ctx context.Context, username string) (*dao.User, error) {
	client := dbc.DB

	users := []dao.User{}

	err := client.WithContext(ctx).Where("username = ?", username).Limit(1).Find(&users).Error
	if len(users) == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := &users[0]

	return user, nil
}
