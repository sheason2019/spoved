package password

import (
	"fmt"

	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func DecodePassword(info *account.AccountInfo) *exception.Exception {
	pwd := crypto_service.DecodeString(info.Password)
	salt := pwd[len(pwd)-len(info.Salt):]
	if salt != info.Salt {
		return exception.New(fmt.Errorf("解密失败，盐不对等 %s %s", salt, info.Salt))
	}

	info.Password = pwd[:len(pwd)-len(info.Salt)]
	return nil
}
