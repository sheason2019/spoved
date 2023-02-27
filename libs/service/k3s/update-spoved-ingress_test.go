package k3s_service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	k3s_service "github.com/sheason2019/spoved/libs/service/k3s"
)

func TestUpdateSpovedIngress(t *testing.T) {
	do := &dao.DeployOrder{
		ServiceName: "test-service",
	}

	ingress, err := k3s_service.UpdateSpovedIngress(context.TODO(), do)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(ingress)
}
