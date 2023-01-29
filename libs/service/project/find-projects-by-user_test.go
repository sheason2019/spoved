package project_service_test

import (
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func TestFindProjectsByUser(t *testing.T) {
	user := ent.User{}
	user.ID = 1

	pg := common.Pagination{
		Page:     1,
		PageSize: 50,
	}

	projs, e := project_service.FindProjectsByUser(&user, &pg)
	if e != nil {
		e.Panic()
	}

	fmt.Println(projs)
}
