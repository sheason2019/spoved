package compile_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
)

func Compile(ctx context.Context, image, nextVersion, branch string, proj *dao.Project, usr *dao.User) (*dao.CompileOrder, error) {
	order := &dao.CompileOrder{
		Branch:     branch,
		Image:      image,
		StatusCode: 0,
		Version:    nextVersion,
		Operator:   *usr,
		Project:    *proj,
	}
	err := CreateCompileOrder(ctx, order)

	compileCtx := CompileContext{
		Context:      ctx,
		CompileOrder: order,
	}

	go CompileRun(compileCtx)

	return order, err
}
