package account_service_test

import (
	"testing"

	account_service "github.com/sheason2019/spoved/libs/service/account"
)

func TestGetRsaPair(t *testing.T) {
	_, err := account_service.GetRsaPair()
	if err != nil {
		t.Error(err)
	}
}

func TestInitRsa(t *testing.T) {
	_, err := account_service.InitRsa()
	if err != nil {
		t.Error(err)
	}
}

func TestMustGetRsaPair(t *testing.T) {
	account_service.MustGetRsaPair()
}
