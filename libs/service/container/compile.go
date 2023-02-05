package container_service

import (
	"fmt"

	"github.com/pkg/errors"
	file_service "github.com/sheason2019/spoved/libs/service/file"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

// 编译项目
func Compile(image, version, branch string, projectId int) error {
	// 操作人权限校验
	// 镜像校验
	if !ValidateImage(image) {
		return errors.WithStack(errors.New("不支持的Image"))
	}

	// 获取Project，以获取git url和dir path
	proj, err := project_service.FindProjectById(projectId)
	if err != nil {
		return err
	}
	if proj == nil {
		return errors.WithStack(errors.New("Project不存在"))
	}

	// 搜索已发布的版本号
	lastRecord, err := FindLastRecordByProjectId(projectId)
	if err != nil {
		return err
	}
	var currentVersion string
	if lastRecord == nil {
		currentVersion = "0.0.0"
	} else {
		currentVersion = lastRecord.Version
	}
	nv, err := nextVersion(currentVersion, version)
	if err != nil {
		return err
	}

	// Git代码复制到的地址
	projDir := proj.DirPath + "/" + nv

	// 拉取代码
	output, err := file_service.GitClone(proj.GitURL, projDir, branch)
	if err != nil {
		return err
	}
	fmt.Println(output)

	// 执行编译命令
	output, err = CompileRun(projDir)
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}
