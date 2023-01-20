package password_test

import (
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/idl-lib/account"
	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
	"github.com/sheason2019/spoved/libs/service/account/password"
)

func TestDecodePassword(t *testing.T) {
	pwd := "test password"
	salt := "test salt"
	cipherPwd := crypto_service.EncodeString(pwd + salt)
	fmt.Println("cipherPwd:", cipherPwd)
	info := account.AccountInfo{
		Username: "",
		Password: cipherPwd,
		Salt:     salt,
	}
	e := password.DecodePassword(&info)
	if e != nil {
		t.Error(e.Print())
	}
	if info.Password != pwd {
		t.Errorf("解密后的密码与原密码不相同")
	}
	fmt.Println(info.Password)
}

func TestEncodePassword(t *testing.T) {
	pwd := "test password"
	cipherPwd, salt := password.EncodePassword(pwd)
	fmt.Println(cipherPwd, salt)
}
