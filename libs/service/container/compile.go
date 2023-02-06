package container_service

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/env"
	git_service "github.com/sheason2019/spoved/libs/service/git"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

// 编译项目
func Compile(image, version, branch string, projectId int, username string) (string, error) {
	outputs := []string{}
	// 操作人权限校验
	// 镜像校验
	if !ValidateImage(image) {
		return "", errors.WithStack(errors.New("不支持的Image"))
	}

	// 获取Project，以获取git url和dir path
	proj, err := project_service.FindProjectById(projectId)
	if err != nil {
		return "", err
	}
	if proj == nil {
		return "", errors.WithStack(errors.New("Project不存在"))
	}

	// 搜索已发布的版本号
	lastRecord, err := FindLastRecordByProjectId(projectId)
	if err != nil {
		return "", err
	}
	var currentVersion string
	if lastRecord == nil {
		currentVersion = "0.0.0"
	} else {
		currentVersion = lastRecord.Version
	}
	nv, err := nextVersion(currentVersion, version)
	if err != nil {
		return "", err
	}

	// Git代码复制到的地址
	projDir := env.DataRoot + proj.DirPath + "/" + nv

	// 拉取代码
	output, err := git_service.GitClone(proj.GitURL, projDir, branch, username)
	if err != nil {
		return "", err
	}
	outputs = append(outputs, output)

	// 执行编译命令
	output, err = CompileRun(context.Background(), projDir)
	if err != nil {
		return output, err
	}
	outputs = append(outputs, output)

	return strings.Join(outputs, "\n"), nil
}
