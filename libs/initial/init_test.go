package initial

import (
	"context"
	"fmt"
	"testing"
)

func TestInitial(t *testing.T) {
	Initial(context.TODO())
}

func TestCreateRootUser(t *testing.T) {
	// 初始化根用户后在Stdout中展示根用户密码
	defer ShowPassword()

	root, err := createRootUser(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(root)
}

func TestInitRootUser(t *testing.T) {
	root, err := initRootUser(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(root)
}
