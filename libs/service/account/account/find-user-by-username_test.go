package account_service_test

import (
	"context"
	"fmt"
	"testing"

	account_service "github.com/sheason2019/spoved/libs/service/account/account"
)

func TestExist(t *testing.T) {
	name := "root"
	usr, e := account_service.FindUserByUsername(context.TODO(), name)
	if e != nil {
		t.Errorf("%+v", e)
	}
	fmt.Println(usr)
}

func TestNotExist(t *testing.T) {
	name := "sheason2020"
	usr, e := account_service.FindUserByUsername(context.TODO(), name)
	if usr == nil {
		return
	}

	t.Errorf("%+v", e)
}
