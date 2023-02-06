// 文件服务，所有文件操作都需要使用这个service
package file_service

import (
	"os"
	"strings"

	"github.com/sheason2019/spoved/libs/env"
)

func Read(path string) (string, error) {
	c, e := os.ReadFile(env.DataRoot + path)
	return string(c), e
}

func Write(c, path string) error {
	exist, err := Exist(path)
	if err != nil {
		return err
	}

	path = env.DataRoot + path
	if !exist {
		os.MkdirAll(path[:strings.LastIndex(path, "/")], os.ModePerm)
	}

	return os.WriteFile(path, []byte(c), 0666)
}

func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func Mkdir(path string) {
	path = env.DataRoot + path

	os.MkdirAll(path, os.ModePerm)
}
