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

func initSpovedIngress(ctx context.Context, root *dao.User) error {
	// 初始化Spoved Project
	proj, err := createSpovedIngressProject(ctx, root)
	if err != nil {
		return errors.WithStack(err)
	}

	version, err := compile_service.FindNextVersionForProject(ctx, int(proj.ID), "Patch")
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建CompileOrder
	co := &dao.CompileOrder{
		Image:      "golang:1.20.0-alpine3.17",
		Version:    version,
		StatusCode: 0,
		Branch:     init_branch,
		Env: map[string]string{
			"PRODUCTION": "true",
			"BUILD_TYPE": "SPOVED_INGRESS",
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
		StatusCode:   0,
		Operator:     *root,
		CompileOrder: *co,
	}

	err = deploy_service.CreateDeployOrder(ctx, do)
	if err != nil {
		return errors.WithStack(err)
	}

	err = deploy_service.DeployRun(ctx, do)
	if err != nil {
		return errors.WithStack(err)
	}

	// 更新Ingress
	_, err = k3s_service.UpdateSpovedIngress(ctx, do)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createSpovedIngressProject(ctx context.Context, root *dao.User) (*dao.Project, error) {
	// 如果Spoved已存在则跳过创建步骤
	proj, err := project_service.FindProject(ctx, root.Username, "spoved")
	if err != nil {
		return nil, err
	}
	if proj != nil {
		return proj, nil
	}

	return project_service.CreateProject(ctx, &project.Project{
		ProjectName: "spoved-ingress",
		GitUrl:      "https://github.com/sheason2019/spoved",
	}, root)
}
