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

	// 创建初始化时伪造的CompileOrder
	co, err := createSpovedCompileOrder(ctx, root, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建初始化时伪造的DeployOrder
	_, err = createSpovedDeployOrder(ctx, root, co)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
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

// 创建伪造的 Spoved CompileOrder
func createSpovedCompileOrder(ctx context.Context, root *dao.User, proj *dao.Project) (*dao.CompileOrder, error) {
	// 如果已经存在CompileOrder，则表示目前是自举部署，这种情况下不必再伪造CompileOrder
	lastCo, err := compile_service.FindLastOrderByProjectId(ctx, int(proj.ID))
	if err != nil {
		return nil, err
	}
	if lastCo != nil {
		return lastCo, nil
	}

	order := &dao.CompileOrder{
		Image:      "golang:1.20.0-alpine3.17",
		Version:    "0.0.1",
		StatusCode: 1,
		Branch:     "master",

		Operator: *root,
		Project:  *proj,
	}

	err = compile_service.CreateCompileOrder(ctx, order)

	return order, err
}

// 创建伪造的 Spoved DeployOrder
func createSpovedDeployOrder(ctx context.Context, root *dao.User, co *dao.CompileOrder) (*dao.DeployOrder, error) {
	// 如果已经存在DeployOrder，则表示目前是自举部署，这种情况下不必再伪造DeployOrder
	lastDo, err := deploy_service.FindLastOrderByCompileOrderID(ctx, int(co.ID))
	if err != nil {
		return nil, err
	}
	if lastDo != nil {
		return lastDo, nil
	}

	do := &dao.DeployOrder{
		Image:        co.OutImageName(),
		StatusCode:   1,
		Operator:     *root,
		CompileOrder: *co,
	}

	err = deploy_service.CreateDeployOrder(ctx, do)

	return do, err
}
