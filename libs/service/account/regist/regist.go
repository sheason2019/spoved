package regist_service

import (
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	"github.com/sheason2019/spoved/libs/service/account/password"
)

func Regist(accountInfo *account.AccountInfo) (*ent.User, error) {
	// 将用户上传的PWD解析成明文
	password.DecodePassword(accountInfo)
	// 校验用户的输入信息是否合规
	e := RegistValidate(accountInfo)
	if e != nil {
		return nil, e
	}
	// 然后将用户输入的密码转换为Salt+Hash
	cipherPwd, salt := password.EncodePassword(accountInfo.Password)
	// 创建用户
	usr, e := CreateUser(accountInfo.Username, cipherPwd, salt)
	if e != nil {
		return nil, e
	}

	return usr, nil
}
