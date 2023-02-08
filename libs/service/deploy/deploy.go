package deploy_service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/dbc"
)

func Deploy(ctx context.Context, operator *ent.User, record *ent.CompileRecord, image string) (any, error) {
	client := dbc.GetClient()

	// 创建部署工单
	deploy, err := client.DeployRecord.Create().
		SetCreatedAt(time.Now()).
		SetImage(image).
		SetStatusCode(0).
		SetContainerHash("").
		AddCompileRecord(record).
		AddOperator(operator).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 执行部署逻辑

	return deploy, nil
}
