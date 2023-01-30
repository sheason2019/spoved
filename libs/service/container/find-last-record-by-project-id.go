package container_service

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindLastRecordByProjectId(id int) (*ent.CompileRecord, *exception.Exception) {
	record, err := dbc.GetClient().CompileRecord.Query().
		Where(
			compilerecord.HasProjectWith(
				project.IDEQ(id),
			),
		).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, exception.New(err)
	}

	return record, nil
}
