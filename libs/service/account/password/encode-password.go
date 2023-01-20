package password

import (
	"crypto/md5"
	"fmt"

	"github.com/sheason2019/spoved/libs/utils"
)

func EncodePassword(password string) (cipherPassword, salt string) {
	salt = utils.RandomStr(32)
	cipherPassword = fmt.Sprintf("%x", md5.Sum([]byte(password+salt)))

	return
}
