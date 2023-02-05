package account_service_test

import (
	"fmt"
	"testing"

	account_service "github.com/sheason2019/spoved/libs/service/account/account"
)

func TestExist(t *testing.T) {
	name := "sheason"
	usr, e := account_service.FindUserByUsername(name)
	if e != nil {
		t.Errorf("%+v", e)
	}
	fmt.Println(usr)
}

func TestNotExist(t *testing.T) {
	name := "sheason2020"
	usr, e := account_service.FindUserByUsername(name)
	if usr == nil {
		return
	}

	t.Errorf("%+v", e)
}
