package k3s_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	networking_v1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 为Spoved Service创建Ingress
func CreateSpovedIngress(ctx context.Context, proj *dao.Project) (*networking_v1.Ingress, error) {
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
