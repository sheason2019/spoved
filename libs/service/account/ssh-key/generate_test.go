package ssh_key_test

import (
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/dao"
	ssh_key "github.com/sheason2019/spoved/libs/service/account/ssh-key"
)

var usr = dao.User{
	Username: "sheason",
}

func TestGenerateKey(t *testing.T) {

	err := ssh_key.GenerateSshKey(&usr)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetSshKey(t *testing.T) {
	pair, err := ssh_key.GetSshKey(&usr)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", pair)
}

func TestGetSshKeyForce(t *testing.T) {
	pair, err := ssh_key.GetSshKeyForce(&usr)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", pair)
}
