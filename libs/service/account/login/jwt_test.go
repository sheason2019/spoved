package login_service_test

import (
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	login_service "github.com/sheason2019/spoved/libs/service/account/login"
)

func TestJwt(t *testing.T) {
	usr := dao.User{
		Username:     "sheason",
		PasswordHash: "password_hash",
		PasswordSalt: "password_salt",
	}

	token, e := login_service.GenerateJwt(&usr)
	if e != nil {
		t.Errorf("%+v", e)
	}

	fmt.Println("token", token)

	claims, e := login_service.ParseJwt(token)
	if e != nil {
		t.Errorf("%+v", e)
	}

	fmt.Println(claims.User)
}
