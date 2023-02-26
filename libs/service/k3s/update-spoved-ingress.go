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
	ingress, err := clientSet.NetworkingV1().Ingresses("default").Get(ctx, "spoved-ingress", v1.GetOptions{})
	if err != nil {
		if !k8s_err.IsNotFound(err) {
			return nil, errors.WithStack(err)
		}
	} else {
		return ingress, nil
	}

	ingress = &networking_v1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name: "spoved-ingress",
			Annotations: map[string]string{
				"ingress.kubernetes.io/ssl-redirect": "false",
			},
		},
		Spec: networking_v1.IngressSpec{
			DefaultBackend: &networking_v1.IngressBackend{
				Service: &networking_v1.IngressServiceBackend{
					Name: do.ServiceName,
					Port: networking_v1.ServiceBackendPort{
						Number: 80,
					},
				},
			},
		},
	}

	return clientSet.NetworkingV1().Ingresses("default").Update(ctx, ingress, v1.UpdateOptions{})
}
