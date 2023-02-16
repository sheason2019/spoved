package compile_service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func nextVersion(version string, variant string) (string, error) {
	strList := strings.Split(version, ".")
	if len(strList) != 3 {
		return "", errors.WithStack(errors.New("版本号数据有误，请检查数据库"))
	}
	versionList := make([]int, 3)
	for i, v := range strList {
		value, err := strconv.Atoi(v)
		if err != nil {
			return "", errors.WithStack(errors.New("Value转换失败" + v))
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
		return "", errors.WithStack(errors.New("Variant参数只能为Patch、Minor或Major"))
	}

	return fmt.Sprintf("%d.%d.%d", versionList[0], versionList[1], versionList[2]), nil
}

func FindNextVersionForProject(ctx context.Context, projectId int, variant string) (string, error) {
	lastRecord, err := FindLastOrderByProjectId(ctx, projectId)
	if err != nil {
		return "", errors.WithStack(err)
	}

	currentVersion := "0.0.0"
	if lastRecord != nil {
		currentVersion = lastRecord.Version
	}

	nv, err := nextVersion(currentVersion, variant)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return nv, nil
}
