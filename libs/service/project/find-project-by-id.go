package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindProjectById(ctx context.Context, id int) (*dao.Project, error) {
	client := dbc.DB

	projDao := &dao.Project{}
	err := client.
		WithContext(ctx).
		Preload("Creator").
		Where("id = ?", id).
		Find(projDao).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return projDao, nil
}
