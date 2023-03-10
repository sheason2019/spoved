package initial

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func initSpoved(ctx context.Context, root *dao.User) error {
	// 初始化Spoved Project
	proj, err := createSpovedProject(ctx, root)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建与Project绑定的Service
	err = k3s_service.CreateServiceByProject(ctx, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建Ingress
	_, err = k3s_service.CreateSpovedIngress(ctx, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建CompileOrder
	// 否则创建编译工单并执行编译
	co := &dao.CompileOrder{
		Image:      "golang:1.20.0-alpine3.17",
		Version:    "0.0.1",
		StatusCode: 1,
		Branch:     "feat/auto-build",
		Env: map[string]string{
			"PRODUCT":    "true",
			"BUILD_TYPE": "SPOVED",
		},

		Operator: *root,
		Project:  *proj,
	}

	err = compile_service.CreateCompileOrder(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compile_service.CompileRun(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建DeployOrder，并在k3s中构建Deployment
	do := &dao.DeployOrder{
		Image:        "golang:1.20.0-alpine3.17",
		StatusCode:   1,
		Operator:     *root,
		CompileOrder: *co,
	}

	err = deploy_service.CreateDeployOrder(ctx, do)
	if err != nil {
		return errors.WithStack(err)
	}

	return deploy_service.DeployRun(ctx, do)
}

func createSpovedProject(ctx context.Context, root *dao.User) (*dao.Project, error) {
	// 如果Spoved已存在则跳过创建步骤
	proj, err := project_service.FindProject(ctx, root.Username, "spoved")
	if err != nil {
		return nil, err
	}
	if proj != nil {
		return proj, nil
	}

	return project_service.CreateProject(ctx, &project.Project{
		ProjectName: "spoved",
		GitUrl:      "https://github.com/sheason2019/spoved",
	}, root)
}
