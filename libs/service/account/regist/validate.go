package regist_service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/idl-lib/account"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
)

// 检测用户的注册信息是否合法
func RegistValidate(ctx context.Context, accountInfo *account.AccountInfo) error {
	// 检查用户名是否重复
	usr, e := account_service.FindUserByUsername(ctx, accountInfo.Username)
	if e != nil {
		return e
	}
	if usr != nil {
		return errors.WithStack(fmt.Errorf("用户名 %s 已存在", accountInfo.Username))
	}
	// 检测用户名是否符合规则
	e = validateUsername(accountInfo.Username)
	if e != nil {
		return e
	}

	// 检测密码长度
	e = validatePassword(accountInfo.Password)
	if e != nil {
		return e
	}

	return nil
}

// 用户名规则
func validateUsername(name string) error {
	// 前端URL关键字
	keywords := []string{"login", "regist", "profile", "service", "new"}
	for _, keyword := range keywords {
		if keyword == name {
			return errors.WithStack(errors.New("使用了URL关键字作为用户名"))
		}
	}

	// 命名规则
	match := ruleUsername(name)
	if !match {
		return errors.WithStack(errors.New("不符合命名规则，仅能使用字母、数字、下划线以及-符号组成用户名"))
	}

	return nil
}

// 使用正则表达式约束用户名
func ruleUsername(name string) bool {
	nameReg, _ := regexp.Compile(`^[\w\-_]{1,24}$`)
	return nameReg.Match([]byte(name))
}

// 检测用户密码
func validatePassword(pwd string) error {
	// 密码长度4-48位
	length := len(pwd)
	if length < 4 || length > 48 {
		return errors.WithStack(errors.New("密码长度必须在4~48位之间"))
	}
	return nil
}
