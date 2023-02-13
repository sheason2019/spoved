package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	"gorm.io/gorm"
)

func FindCompileRecords(ctx context.Context, projectId int, pagination *common.Pagination) ([]dao.CompileOrder, int, error) {
	client := dbc.GetClient()

	records := make([]dao.CompileOrder, 0)
	err := client.WithContext(ctx).
		Model(&records).
		InnerJoins("Project", client.Where(&dao.Project{Model: gorm.Model{ID: uint(projectId)}})).
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
		InnerJoins("Project", client.Where(&dao.Project{Model: gorm.Model{ID: uint(projectId)}})).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return records, int(count), nil
}
