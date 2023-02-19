package project_service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func TestFindProjectsByUser(t *testing.T) {
	user := dao.User{}
	user.ID = 1

	pg := common.Pagination{
		Page:     1,
		PageSize: 50,
	}

	projs, err := project_service.FindProjectsByUser(context.TODO(), &user, &pg)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	fmt.Println(projs)
}
