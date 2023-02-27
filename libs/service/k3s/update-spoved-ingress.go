package k3s_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	networking_v1 "k8s.io/api/networking/v1"
	k8s_err "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 为Spoved Service创建Ingress
func UpdateSpovedIngress(ctx context.Context, do *dao.DeployOrder) (*networking_v1.Ingress, error) {
	pathType := networking_v1.PathTypePrefix

	exist, err := spovedIngressExist(ctx)
	if err != nil {
		return nil, err
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
											Name: do.ServiceName,
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

	if exist {
		return clientSet.NetworkingV1().Ingresses("default").Update(ctx, ingress, v1.UpdateOptions{})
	} else {
		return clientSet.NetworkingV1().Ingresses("default").Create(ctx, ingress, v1.CreateOptions{})
	}
}

func spovedIngressExist(ctx context.Context) (bool, error) {
	// 尝试获取spoved-ingress
	_, err := clientSet.NetworkingV1().Ingresses("default").Get(ctx, "spoved-ingress", v1.GetOptions{})
	if err != nil {
		// 发生错误时，如果不是没有找到指定的ingress，则抛出错误
		if !k8s_err.IsNotFound(err) {
			return false, errors.WithStack(err)
		} else {
			return false, nil
		}
	} else {
		// 如果找到了指定的ingress，则返回true
		return true, nil
	}
}
