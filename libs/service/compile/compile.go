package compile_service

import (
	"context"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
)

func Compile(ctx context.Context, image, nextVersion, branch string, proj *dao.Project, usr *dao.User) (*dao.CompileOrder, error) {
	statusCode := 0

	order := &dao.CompileOrder{
		Branch:     branch,
		Image:      image,
		StatusCode: statusCode,
		Version:    nextVersion,
		Operator:   *usr,
		Project:    *proj,
	}
	err := CreateCompileOrder(ctx, order)

	go func() {
		_, err := CompileRun(image, nextVersion, branch, proj, usr.Username)

		if err != nil {
			order.StatusCode = -1
		} else {
			order.StatusCode = 1
		}
		dbc.GetClient().Save(&order)
	}()

	return order, err
}
