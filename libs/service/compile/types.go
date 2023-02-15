package compile_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
)

type CompileContext struct {
	context.Context
	CompileOrder *dao.CompileOrder
}
