package container_service

import (
	"errors"
	"fmt"

	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/compile"
	file_service "github.com/sheason2019/spoved/libs/service/file"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

// 编译项目
func Compile(payload *compile.CompileRecord) (e *exception.Exception) {
	// 操作人权限校验
	// 镜像校验
	if !ValidateImage(payload.Image) {
		return exception.New(errors.New("不支持的Image"))
	}

	// 获取Project，以获取git url和dir path
	proj, e := project_service.FindProjectById(payload.ProjectId)
	if e != nil {
		return e.Wrap()
	}
	if proj == nil {
		return exception.New(errors.New("Project不存在"))
	}

	// 搜索已发布的版本号
	lastRecord, e := FindLastRecordByProjectId(payload.ProjectId)
	if e != nil {
		return e.Wrap()
	}
	nv, e := nextVersion(lastRecord.Version, payload.Version)
	if e != nil {
		return e.Wrap()
	}

	// 拉取代码
	output, e := file_service.GitClone(proj.GitURL, proj.DirPath+"/"+nv, payload.Branch)
	if e != nil {
		return e.Wrap()
	}
	fmt.Println(output)

	// 执行编译命令

	return nil
}
