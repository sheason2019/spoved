package k3s_service

import (
	"context"
	"flag"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 初始化前端Pod
func InitSpovedFe(ctx context.Context) error {
	// 寻找 spoved-fe 的svc
	_, err := findSpovedFeService(ctx, clientSet)
	// 如果 spoved-fe 已初始化则不重复创建
	if err == nil {
		fmt.Println("spoved-fe 已存在，跳过初始化步骤。")
		return nil
	}

	// 若发生异常错误则上抛错误
	if !isNotFound(err) {
		return errors.WithStack(err)
	}

	// 否则创建 spoved-fe
	_, _, err = createSpovedFe(ctx, clientSet)
	if err != nil {
		return err
	}

	return nil
}

func getConfig() (*rest.Config, error) {
	// Pod内获取Config
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	// Pod外获取Config
	kubeconfig := flag.String("kubeconfig", "/etc/rancher/k3s/k3s.yaml", "path to the kubeconfig file")
	flag.Parse()

	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}

// 寻找spoved-fe-service服务，判断是否已经启动
func findSpovedFeService(ctx context.Context, clientset *kubernetes.Clientset) (*v1.Service, error) {
	return clientset.CoreV1().Services("default").Get(ctx, "root--spoved-fe-service", meta_v1.GetOptions{})
}

// 创建 Spoved Fe 的 Service 和 Deployment
func createSpovedFe(ctx context.Context, clientset *kubernetes.Clientset) (deployment *appv1.Deployment, service *v1.Service, err error) {
	deployment, err = createSpovedFeDeployment(ctx, clientset)
	if err != nil {
		return
	}
	service, err = createSpovedFeService(ctx, clientset)
	if err != nil {
		return
	}

	return
}

// 创建 Spoved Fe 的 Service
func createSpovedFeService(ctx context.Context, clientset *kubernetes.Clientset) (*v1.Service, error) {
	service := &v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "root--spoved-fe-service",
			Namespace: "default",
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Port:     80,
					Protocol: "TCP",
				},
			},
			Selector: map[string]string{
				"run": "spoved-fe-deployment",
			},
		},
	}

	return clientset.CoreV1().Services("default").Create(ctx, service, meta_v1.CreateOptions{})
}

// 创建 Spoved Fe 的 Deployment
func createSpovedFeDeployment(ctx context.Context, clientset *kubernetes.Clientset) (*appv1.Deployment, error) {
	deployment := &appv1.Deployment{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: "deployment-",
		},
		Spec: appv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &meta_v1.LabelSelector{
				MatchLabels: map[string]string{
					"owner": "root",
					"name":  "spoved-fe",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: meta_v1.ObjectMeta{
					Labels: map[string]string{
						"run": "spoved-fe-deployment",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "spoved-fe",
							Image: "sheason/spoved-fe",
							Ports: []v1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	return clientset.AppsV1().Deployments("default").Create(ctx, deployment, meta_v1.CreateOptions{})
}

func isNotFound(err error) bool {
	reg := regexp.MustCompile(`services \".+\" not found`)
	return reg.MatchString(err.Error())
}

func int32Ptr(i int32) *int32 { return &i }
