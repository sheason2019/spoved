package ssh_key

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/utils"
)

type SshKeyPair struct {
	PublicKey  string
	PrivateKey string
}

// 为指定用户创建ssh key
func GenerateSshKey(usr *ent.User) error {
	dirPath := utils.GetRootPath() + "/data/account/" + usr.Username + "/.ssh"
	os.MkdirAll(dirPath, os.ModePerm)

	var out bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", "ssh-keygen -N \"\" -f "+dirPath+"/id_rsa <<< y")
	cmd.Stderr = &out
	cmd.Stdout = &out

	configString := "StrictHostKeyChecking no\nUserKnownHostsFile /dev/null"
	os.WriteFile(dirPath+"/config", []byte(configString), os.ModePerm)

	err := cmd.Run()
	if err != nil {
		fmt.Println(out.String())
	}
	return err
}

// 获取指定用户的ssh key
func GetSshKey(usr *ent.User) (*SshKeyPair, error) {
	dirPath := utils.GetRootPath() + "/data/account/" + usr.Username + "/.ssh"
	priv, err := readFile(dirPath + "/id_rsa")
	if err != nil {
		return nil, err
	}

	pub, err := readFile(dirPath + "/id_rsa.pub")
	if err != nil {
		return nil, err
	}

	return &SshKeyPair{
		PrivateKey: priv,
		PublicKey:  pub,
	}, nil
}

// 获取指定用户的ssh key，如果不存在则创建
func GetSshKeyForce(usr *ent.User) (*SshKeyPair, error) {
	pair, err := GetSshKey(usr)
	if err != nil {
		return nil, err
	}
	// 若公钥或私钥不存在，则重新创建
	if pair.PrivateKey == "" || pair.PublicKey == "" {
		err = GenerateSshKey(usr)
		if err != nil {
			return nil, err
		}
		// 递归保证一定获取
		return GetSshKeyForce(usr)
	} else {
		return pair, err
	}
}

// 读取文件内容
func readFile(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
