package compile_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateCompileOrder(ctx context.Context, order *dao.CompileOrder) error {
	return dbc.DB.WithContext(ctx).Save(order).Error
}
