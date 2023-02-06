package compile_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/env"
	git_service "github.com/sheason2019/spoved/libs/service/git"
	images "github.com/sheason2019/spoved/libs/service/images"
)

// 编译项目
func CompileRun(image, nextVersion, branch string, proj *ent.Project, username string) ([]string, error) {
	outputs := []string{}

	// TODO: 操作人权限校验

	// 镜像校验
	outputs = append(outputs, "正在校验镜像信息："+image)
	if !images.ValidateImage(image) {
		errStr := "不支持的镜像：" + image
		outputs = append(outputs, errStr)
		return outputs, errors.WithStack(errors.New(errStr))
	}

	// Git代码复制到的地址
	projDir := env.DataRoot + proj.DirPath + "/" + nextVersion

	// 拉取代码
	outputs = append(outputs, "正在尝试拉取Git仓库")
	output, err := git_service.GitClone(proj.GitURL, projDir, branch, username)
	if err != nil {
		outputs = append(outputs, "拉取Git仓库失败")
		return outputs, errors.WithStack(err)
	}
	outputs = append(outputs, output)
	outputs = append(outputs, "拉取Git仓库成功")

	// 执行编译命令
	outputs = append(outputs, "正在尝试编译项目")
	output, err = CompileRunBuild(context.Background(), projDir)
	if err != nil {
		outputs = append(outputs, "编译项目失败")
		return outputs, errors.WithStack(err)
	}
	outputs = append(outputs, output)
	outputs = append(outputs, "编译项目成功")

	return outputs, nil
}
