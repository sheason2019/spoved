package transfer

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
)

func CompileRecordToIdl(ctx context.Context, cr *ent.CompileRecord) *compile.CompileRecord {
	record := compile.CompileRecord{}

	record.Id = cr.ID
	record.Branch = cr.Branch
	record.CreateAt = int(cr.CreatedAt.Unix())
	record.Image = cr.Image

	usr := cr.QueryOperator().OnlyX(ctx)
	record.Operator = usr.Username

	record.Output = cr.Output
	proj := cr.QueryProject().OnlyX(ctx)

	record.ProjectId = proj.ID
	record.StatusCode = cr.StatusCode
	record.Version = cr.Version

	return &record
}
