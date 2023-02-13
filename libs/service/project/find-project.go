package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindProject(ctx context.Context, username, projName string) (*dao.Project, error) {
	client := dbc.GetClient()

	projDao := &dao.Project{}
	err := client.WithContext(ctx).
		Where("project_name = ?", projName).
		InnerJoins("Users", client.Where(
			&dao.User{Username: username}),
		).
		Limit(0).
		Find(projDao).
		Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.WithStack(errors.New("未找到指定的项目" + `/` + username + `/` + projName))
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return projDao, nil
}
