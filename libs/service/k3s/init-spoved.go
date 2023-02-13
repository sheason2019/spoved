package k3s_service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	deploy_service "github.com/sheason2019/spoved/libs/service/deploy"
	project_service "github.com/sheason2019/spoved/libs/service/project"
	"github.com/sheason2019/spoved/libs/utils"

	networking_v1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func InitSpoved() error {
	ctx := context.TODO()
	// 初始化Root用户，密码为32位随机字符串
	root, err := createRootUser(ctx)
	if err != nil {
		return err
	}

	// 初始化Spoved Project
	proj, err := createProject(ctx, root)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建与Project绑定的Service
	err = CreateServiceByProject(ctx, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建Ingress
	_, err = createIngress(ctx, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建初始化时伪造的CompileOrder
	co, err := createCompileOrder(ctx, root, proj)
	if err != nil {
		return errors.WithStack(err)
	}

	// 创建初始化时伪造的DeployOrder
	_, err = createDeployOrder(ctx, root, co)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func createRootUser(ctx context.Context) (*dao.User, error) {
	pwd := utils.RandomStr(32)
	rootUser, err := regist_service.Regist(ctx, "root", pwd)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fmt.Printf("root user created,password: %s\n", pwd)
	return rootUser, nil
}

func createProject(ctx context.Context, root *dao.User) (*dao.Project, error) {
	return project_service.CreateProject(ctx, &project.Project{
		ProjectName: "spoved",
		GitUrl:      "https://github.com/sheason2019/spoved",
	}, root)
}

// 创建伪造的 Spoved CompileOrder
func createCompileOrder(ctx context.Context, root *dao.User, proj *dao.Project) (*dao.CompileOrder, error) {
	order := &dao.CompileOrder{
		Image:      "golang:1.20.0-alpine3.17",
		Version:    "0.0.1",
		StatusCode: 1,
		Branch:     "master",

		Operator: *root,
		Project:  *proj,
	}

	err := compile_service.CreateCompileOrder(ctx, order)

	return order, err
}

// 创建伪造的 Spoved DeployOrder
func createDeployOrder(ctx context.Context, root *dao.User, co *dao.CompileOrder) (*dao.DeployOrder, error) {
	do := &dao.DeployOrder{
		Image:        co.OutImageName(),
		StatusCode:   1,
		Operator:     *root,
		CompileOrder: *co,
	}

	err := deploy_service.CreateDeployOrder(ctx, do)

	return do, err
}

// 为Spoved Service创建Ingress
func createIngress(ctx context.Context, proj *dao.Project) (*networking_v1.Ingress, error) {
	pathType := networking_v1.PathTypePrefix

	exist, err := clientSet.NetworkingV1().Ingresses("default").Get(ctx, "spoved-ingress", v1.GetOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if exist != nil {
		return exist, nil
	}

	ingress := &networking_v1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name: "spoved-ingress",
			Annotations: map[string]string{
				"ingress.kubernetes.io/ssl-redirect": "false",
			},
		},
		Spec: networking_v1.IngressSpec{
			Rules: []networking_v1.IngressRule{
				{
					IngressRuleValue: networking_v1.IngressRuleValue{
						HTTP: &networking_v1.HTTPIngressRuleValue{
							Paths: []networking_v1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networking_v1.IngressBackend{
										Service: &networking_v1.IngressServiceBackend{
											Name: proj.ServiceName,
											Port: networking_v1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return clientSet.NetworkingV1().Ingresses("default").Create(ctx, ingress, v1.CreateOptions{})
}
