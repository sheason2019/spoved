// 文件服务，所有文件操作都需要使用这个service
package file_service

import (
	"os"
	"strings"
)

const path_root = "data/"

func Read(path string) (string, error) {
	c, e := os.ReadFile(path_root + path)
	return string(c), e
}

func Write(c, path string) error {
	path = path_root + path

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path[:strings.LastIndex(path, "/")], os.ModePerm)
		} else {
			return err
		}
	}

	return os.WriteFile(path, []byte(c), 0666)
}
