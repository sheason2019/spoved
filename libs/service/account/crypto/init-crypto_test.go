package crypto_service_test

import (
	"testing"

	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func TestGetRsaPair(t *testing.T) {
	_, err := crypto_service.GetRsaPair()
	if err != nil {
		t.Error(err)
	}
}

func TestInitRsa(t *testing.T) {
	_, err := crypto_service.InitRsa()
	if err != nil {
		t.Error(err)
	}
}

func TestMustGetRsaPair(t *testing.T) {
	crypto_service.MustGetRsaPair()
}
