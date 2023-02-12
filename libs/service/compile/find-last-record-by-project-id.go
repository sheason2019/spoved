package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindLastRecordByProjectId(id int) (*ent.CompileRecord, error) {
	record, err := dbc.GetClient().CompileRecord.Query().
		Where(
			compilerecord.HasProjectWith(
				project.IDEQ(id),
			),
		).
		Order(ent.Desc(compilerecord.FieldCreatedAt)).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return record, nil
}
