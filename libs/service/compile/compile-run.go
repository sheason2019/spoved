package compile_service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

// 编译项目
func CompileRun(ctx context.Context, co *dao.CompileOrder) (err error) {
	defer func(cause error) {
		if cause == nil {
			co.StatusCode = 1
		} else {
			co.StatusCode = -1
		}

		fmt.Println(err)
		dbc.DB.WithContext(ctx).Save(co)
	}(err)

	// 拉取代码
	fmt.Printf("CompileOrder: %d Version: %s 开始拉取代码\n", co.ID, co.Version)
	err = k3s_service.GitClone(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	// 执行编译命令
	fmt.Printf("CompileOrder: %d Version: %s 开始执行编译\n", co.ID, co.Version)
	err = k3s_service.Build(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
