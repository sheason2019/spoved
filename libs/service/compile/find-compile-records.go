package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
)

func FindCompileRecords(ctx context.Context, projectId int, pagination *common.Pagination) ([]*ent.CompileRecord, int, error) {
	client := dbc.GetClient()

	records, err := client.CompileRecord.Query().
		Limit(pagination.PageSize).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Where(
			compilerecord.HasProjectWith(
				project.IDEQ(projectId),
			),
		).All(ctx)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	count, err := client.CompileRecord.Query().
		Where(
			compilerecord.HasProjectWith(
				project.IDEQ(projectId),
			),
		).
		Count(ctx)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return records, count, nil
}
