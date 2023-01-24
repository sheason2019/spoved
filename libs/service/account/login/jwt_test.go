package login_service_test

import (
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/ent"
	login_service "github.com/sheason2019/spoved/libs/service/account/login"
)

func TestJwt(t *testing.T) {
	usr := ent.User{
		Username:     "sheason",
		PasswordHash: "password_hash",
		PasswordSalt: "password_salt",
	}

	token, e := login_service.GenerateJwt(&usr)
	if e != nil {
		e.Panic()
	}

	fmt.Println("token", token)

	claims, e := login_service.ParseJwt(token)
	if e != nil {
		e.Panic()
	}

	fmt.Println(claims.User)
}
