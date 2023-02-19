package initial

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/env"
	account_service "github.com/sheason2019/spoved/libs/service/account/account"
	regist_service "github.com/sheason2019/spoved/libs/service/account/regist"
	"github.com/sheason2019/spoved/libs/utils"
)

// 初始化根用户
func initRootUser(ctx context.Context) (root *dao.User, err error) {
	// 初始化根用户后在Stdout中展示根用户密码
	defer func() {
		err := showPassword()
		if err != nil {
			panic(fmt.Errorf("get root password failure: %w", err))
		}
	}()

	root, err = account_service.FindUserByUsername(ctx, "root")
	if err != nil {
		return
	}
	if root != nil {
		return
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

	err = os.WriteFile(env.DataRoot+"/ROOT_PASSWORD", []byte(pwd), os.ModePerm)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rootUser, nil
}

func showPassword() error {
	pwd, err := os.ReadFile(env.DataRoot + "/ROOT_PASSWORD")
	if err != nil {
		return err
	}

	fmt.Println("ROOT PASSWORD: " + string(pwd))
	return nil
}
