package k3s_service

import (
	"context"
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGitClone(t *testing.T) {
	ctx := context.TODO()

	job, err := clientSet.BatchV1().Jobs("default").Get(ctx, "git-clone-co-id-6", v1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("job %+v", job.Status.Conditions)
}

func TestCreateIngress(t *testing.T) {
	ingress, err := CreateSpovedIngress(context.TODO(), &dao.Project{})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(ingress)
}
