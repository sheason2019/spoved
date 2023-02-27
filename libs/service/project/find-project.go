package project_service

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindProject(ctx context.Context, username, projName string) (*dao.Project, error) {
	client := dbc.DB

	projDao := &dao.Project{}
	err := client.WithContext(ctx).
		Joins("inner join users on users.id = projects.creator_id").
		Where("project_name = ?", projName).
		Where("users.username = ?", username).
		Preload("Creator").
		Limit(1).
		Find(projDao).
		Error

	if err == gorm.ErrRecordNotFound || projDao.ID == 0 {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("FindProjectError:%w", err)
	}

	return projDao, nil
}
