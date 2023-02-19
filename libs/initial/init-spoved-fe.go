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

// 初始化前端Project
func initSpovedFe(ctx context.Context, root *dao.User) error {
	// 初始化Project
	proj, _, err := createSpovedFeProject(ctx, root)
	if err != nil {
		return errors.WithStack(err)
	}

	// 初始化Spoved-fe 的 Service
	err = k3s_service.CreateServiceByProject(ctx, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建 Compile Order
	co := &dao.CompileOrder{
		Image:    "node:16-alpine",
		Version:  "0.0.1",
		Branch:   "test/build",
		Project:  *proj,
		Operator: *root,
	}
	err = compile_service.CreateCompileOrder(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compile_service.CompileRun(ctx, co)
	if err != nil {
		return errors.WithStack(err)
	}

	// 执行部署逻辑，在k3s中为spoved-fe创建deployment
	do, err := deploy_service.CreateDeployOrderByCO(ctx, root, co, "root/spoved-nginx")
	if err != nil {
		return errors.WithStack(err)
	}

	err = deploy_service.DeployRun(ctx, do)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createSpovedFeProject(ctx context.Context, root *dao.User) (proj *dao.Project, exist bool, err error) {
	// 如果Spoved-fe已存在则跳过创建步骤
	proj, err = project_service.FindProject(ctx, root.Username, "spoved-fe")
	if err != nil {
		return nil, false, err
	}
	if proj != nil {
		return proj, true, nil
	}

	proj, err = project_service.CreateProject(ctx, &project.Project{
		ProjectName: "spoved-fe",
		GitUrl:      "https://github.com/sheason2019/spoved-fe",
	}, root)

	return
}

func createSpovedFeCompileOrder(ctx context.Context, root *dao.User, proj *dao.Project) {}
