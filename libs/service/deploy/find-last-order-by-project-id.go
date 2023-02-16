package deploy_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindLastOrderByCompileOrderID(ctx context.Context, coID int) (*dao.DeployOrder, error) {
	client := dbc.DB

	order := &dao.DeployOrder{}
	err := client.WithContext(ctx).
		Where("compile_order_id = ?", coID).
		Order("created_at desc").
		Limit(1).
		Find(order).
		Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return order, nil
}
