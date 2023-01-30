// 文件服务，所有文件操作都需要使用这个service
package file_service

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/sheason2019/spoved/exceptions/exception"
)

const path_root = "data/"

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

func GitClone(url, dir, branch string) (string, *exception.Exception) {
	dir = path_root + dir

	outputs := []string{}

	output, err := exec.Command("git", "clone", "--depth", "1", url, dir).Output()
	if err != nil {
		fmt.Println("throw")
		return string(output), exception.New(err)
	}
	outputs = append(outputs, string(output))

	if branch != "master" {
		output, err := exec.Command(
			"sh",
			"-c",
			fmt.Sprintf("cd %s\ngit fetch origin %s\ngit checkout %s\ngit pull", dir, branch, branch),
		).Output()
		if err != nil {
			outputs = append(outputs, string(output))
		}
	}

	return strings.Join(outputs, "\n"), nil
}

func GetAbsPath(p string) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}

	p = path_root + p

	return filepath.Abs(p)
}
