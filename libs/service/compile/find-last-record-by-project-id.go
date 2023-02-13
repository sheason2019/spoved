package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"gorm.io/gorm"
)

func FindLastRecordByProjectId(ctx context.Context, id int) (*dao.CompileOrder, error) {
	client := dbc.GetClient()

	order := &dao.CompileOrder{}
	err := client.WithContext(ctx).
		InnerJoins("Project", client.Where(&dao.Project{Model: gorm.Model{ID: order.ID}})).
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
