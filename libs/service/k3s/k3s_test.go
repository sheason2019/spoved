package k3s_service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientSet = k3s_service.GetClientSet()

func TestGitClone(t *testing.T) {
	ctx := context.TODO()

	job, err := clientSet.BatchV1().Jobs("default").Get(ctx, "git-clone-co-id-6", v1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("job %+v", job.Status.Conditions)
}

func TestCreateIngress(t *testing.T) {
	do := &dao.DeployOrder{}

	ingress, err := k3s_service.UpdateSpovedIngress(context.TODO(), do)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(ingress)
}

var spoved = &dao.Project{
	Creator: dao.User{
		Username: "root",
	},
	ProjectName: "spoved",
}

func TestFindProjectDeployments(t *testing.T) {
	deploys, err := k3s_service.FindProjectDeployments(context.TODO(), spoved)
	if err != nil {
		t.Error(err)
		return
	}

	for _, deploy := range deploys.Items {
		fmt.Printf("%+v\n", deploy)
	}
}

func TestFindProjectServices(t *testing.T) {
	services, err := k3s_service.FindProjectServices(context.TODO(), spoved)
	if err != nil {
		t.Error(err)
		return
	}

	for _, svc := range services.Items {
		fmt.Printf("%+v\n", svc)
	}
}
