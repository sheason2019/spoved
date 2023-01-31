package container_service

import (
	"errors"
	"fmt"

	"github.com/sheason2019/spoved/exceptions/exception"
	file_service "github.com/sheason2019/spoved/libs/service/file"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

// 编译项目
func Compile(image, version, branch string, projectId int) (e *exception.Exception) {
	// 操作人权限校验
	// 镜像校验
	if !ValidateImage(image) {
		return exception.New(errors.New("不支持的Image"))
	}

	// 获取Project，以获取git url和dir path
	proj, e := project_service.FindProjectById(projectId)
	if e != nil {
		return e.Wrap()
	}
	if proj == nil {
		return exception.New(errors.New("Project不存在"))
	}

	// 搜索已发布的版本号
	lastRecord, e := FindLastRecordByProjectId(projectId)
	if e != nil {
		return e.Wrap()
	}
	var currentVersion string
	if lastRecord == nil {
		currentVersion = "0.0.0"
	} else {
		currentVersion = lastRecord.Version
	}
	nv, e := nextVersion(currentVersion, version)
	if e != nil {
		return e.Wrap()
	}

	// Git代码复制到的地址
	projDir := proj.DirPath + "/" + nv

	// 拉取代码
	output, e := file_service.GitClone(proj.GitURL, projDir, branch)
	if e != nil {
		return e.Wrap()
	}
	fmt.Println(output)

	// 执行编译命令
	output, e = CompileRun(projDir)
	if e != nil {
		return e.Wrap()
	}
	fmt.Println(output)

	return nil
}
