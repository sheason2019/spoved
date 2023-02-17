package initial

import (
	"context"
	"fmt"
)

func Initial(ctx context.Context) {
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
