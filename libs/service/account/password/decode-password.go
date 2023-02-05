package password

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	crypto_service "github.com/sheason2019/spoved/libs/service/account/crypto"
)

func DecodePassword(info *account.AccountInfo) error {
	pwd, err := crypto_service.DecodeString(info.Password)
	if err != nil {
		return err
	}
	salt := pwd[len(pwd)-len(info.Salt):]
	if salt != info.Salt {
		return errors.WithStack(fmt.Errorf("解密失败，盐不对等 %s %s", salt, info.Salt))
	}

	info.Password = pwd[:len(pwd)-len(info.Salt)]
	return nil
}
