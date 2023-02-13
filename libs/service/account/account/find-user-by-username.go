package account_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindUserByUsername(ctx context.Context, username string) (*dao.User, error) {
	client := dbc.GetClient()

	user := &dao.User{
		Username: username,
	}

	err := client.WithContext(ctx).Where(&user).Limit(1).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
