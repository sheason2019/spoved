package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
)

// 根据用户名搜索项目
func FindProjectsByUser(ctx context.Context, usr *dao.User, pagination *common.Pagination) ([]dao.Project, error) {
	client := dbc.DB
	projDaos := make([]dao.Project, 0)

	err := client.WithContext(ctx).
		Where("creator_id = ?", usr.ID).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit(pagination.PageSize).
		Find(&projDaos).
		Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return projDaos, nil
}
