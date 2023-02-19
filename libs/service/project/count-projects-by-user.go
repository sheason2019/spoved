package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CountProjectsByUser(ctx context.Context, usr *dao.User) (int64, error) {
	client := dbc.DB

	var count int64
	err := client.WithContext(ctx).
		Model(&dao.Project{}).
		Where("CreatorID = ?", usr.ID).
		Count(&count).
		Error

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}
