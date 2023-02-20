package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
)

func FindCompileRecords(ctx context.Context, projectId int, pagination *common.Pagination) ([]dao.CompileOrder, int, error) {
	client := dbc.DB

	records := make([]dao.CompileOrder, 0)
	err := client.WithContext(ctx).
		Model(&records).
		Joins("inner join projects on compile_orders.project_id = projects.id").
		Where("projects.id = ?", projectId).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit(pagination.PageSize).
		Find(&records).
		Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	var count int64
	err = client.WithContext(ctx).
		Model(&records).
		Joins("inner join projects on compile_orders.project_id = projects.id").
		Where("projects.id = ?", projectId).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return records, int(count), nil
}

func FindLastOrderByProjectId(ctx context.Context, id int) (*dao.CompileOrder, error) {
	client := dbc.DB

	orders := make([]dao.CompileOrder, 0)
	err := client.WithContext(ctx).
		Joins("inner join projects on compile_orders.project_id = projects.id").
		Where("projects.id = ?", id).
		Order("compile_orders.created_at desc").
		Limit(1).
		Find(&orders).
		Error

	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(orders) == 0 {
		return nil, nil
	}

	return &orders[0], nil
}
