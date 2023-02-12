package password

import (
	"crypto/md5"
	"fmt"

	"github.com/sheason2019/spoved/libs/utils"
)

func EncodePassword(password string) (cipherPassword, salt string) {
	salt = utils.RandomStr(32)
	cipherPassword = StringHash(password + salt)

	return
}

func StringHash(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
