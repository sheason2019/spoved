// 文件服务，所有文件操作都需要使用这个service
package file_service

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sheason2019/spoved/libs/env"
	"github.com/sheason2019/spoved/libs/utils"
)

var path_root string

func init() {
	if env.IS_PRODUCT {
		path_root = "data"
	} else {
		path_root = utils.GetRootPath() + "/data"
	}
}

func Read(path string) (string, error) {
	c, e := os.ReadFile(path_root + path)
	return string(c), e
}

func Write(c, path string) error {
	exist, err := Exist(path)
	if err != nil {
		return err
	}

	path = path_root + path
	if !exist {
		os.MkdirAll(path[:strings.LastIndex(path, "/")], os.ModePerm)
	}

	return os.WriteFile(path, []byte(c), 0666)
}

func Exist(path string) (bool, error) {
	path = path_root + path

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
	path = path_root + path

	os.MkdirAll(path, os.ModePerm)
}

func GetAbsPath(p string) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}

	p = path_root + p

	return filepath.Abs(p)
}
