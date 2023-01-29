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
		e.Panic()
	}
	fmt.Println(usr)
}

func TestNotExist(t *testing.T) {
	name := "sheason2020"
	usr, e := account_service.FindUserByUsername(name)
	if usr == nil {
		return
	}

	e.Panic()
}
