package transfer

import (
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
)

func CompileOrderToIdl(co *dao.CompileOrder) *compile.CompileOrder {
	record := compile.CompileOrder{}

	record.Id = int(co.ID)
	record.Branch = co.Branch
	record.CreateAt = int(co.CreatedAt.Unix())
	record.Image = co.Image

	usr := co.Operator
	record.Operator = usr.Username

	proj := co.Project

	record.ProjectId = int(proj.ID)
	record.StatusCode = co.StatusCode
	record.Version = co.Version

	record.Production = co.Production

	return &record
}
