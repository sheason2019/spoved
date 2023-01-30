package container_service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sheason2019/spoved/exceptions/exception"
)

func nextVersion(version string, variant string) (string, *exception.Exception) {
	strList := strings.Split(version, ".")
	if len(strList) != 3 {
		return "", exception.New(errors.New("版本号数据有误，请检查数据库"))
	}
	versionList := make([]int, 3)
	for i, v := range strList {
		value, err := strconv.Atoi(v)
		if err != nil {
			return "", exception.New(errors.New("Value转换失败" + v))
		}
		versionList[i] = value
	}

	if variant == "Patch" {
		versionList[2] = versionList[2] + 1
	} else if variant == "Minor" {
		versionList[1] = versionList[1] + 1
		versionList[2] = 0
	} else if variant == "Major" {
		versionList[0] = versionList[0] + 1
		versionList[1] = 0
		versionList[2] = 0
	} else {
		return "", exception.New(errors.New("Variant参数只能为Patch、Minor或Major"))
	}

	return fmt.Sprintf("%d.%d.%d", versionList[0], versionList[1], versionList[2]), nil
}
