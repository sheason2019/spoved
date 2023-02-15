package initial

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
	"github.com/sheason2019/spoved/libs/utils"
)

// 初始化根用户
func initRootUser(ctx context.Context) (*dao.User, error) {
	root, err := account_service.FindUserByUsername(ctx, "root")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if root != nil {
		return root, err
	}
	// 初始化Root用户，密码为32位随机字符串
	return createRootUser(ctx)
}

func createRootUser(ctx context.Context) (*dao.User, error) {
	pwd := utils.RandomStr(32)
	rootUser, err := regist_service.Regist(ctx, "root", pwd)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fmt.Printf("root user created,password: %s\n", pwd)
	return rootUser, nil
}
