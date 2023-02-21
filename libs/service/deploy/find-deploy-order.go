package deploy_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindDeployOrders(ctx context.Context, projectId int, page, pageSize int) ([]dao.DeployOrder, int, error) {
	client := dbc.DB

	records := make([]dao.DeployOrder, 0)

	err := client.WithContext(context.TODO()).
		Model(&records).
		Joins("CompileOrder").
		Joins("CompileOrder.Project", client.Where("CompileOrder.Project.id = ?", projectId)).
		Preload("Operator").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&records).
		Error

	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	var count int64
	err = client.WithContext(ctx).
		Model(&records).
		Preload("CompileOrder").
		Preload("CompileOrder.Project", client.Where("id = ?", 2)).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return records, int(count), nil
}

func FindLastOrderByCompileOrderID(ctx context.Context, coID int) (*dao.DeployOrder, error) {
	client := dbc.DB

	orders := make([]dao.DeployOrder, 0)
	err := client.WithContext(ctx).
		Where("compile_order_id = ?", coID).
		Order("created_at desc").
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
