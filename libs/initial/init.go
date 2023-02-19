package initial

import (
	"context"
	"fmt"
	"os"

	"github.com/sheason2019/spoved/libs/env"
)

func Initial(ctx context.Context) {
	fmt.Println("正在初始化Data目录")
	err := os.MkdirAll(env.DataRoot, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("正在初始化Root账户")
	root, err := initRootUser(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("正在初始化Spoved服务")
	err = initSpoved(ctx, root)
	if err != nil {
		panic(err)
	}

	fmt.Println("正在初始化Spoved-fe服务")
	err = initSpovedFe(ctx, root)
	if err != nil {
		panic(err)
	}
}
