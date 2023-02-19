package project_service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	project_idl "github.com/sheason2019/spoved/libs/idl-lib/project"
)

func validate(ctx context.Context, proj *project_idl.Project, creator *dao.User) error {
	if len(proj.ProjectName) == 0 {
		return errors.WithStack(errors.New("项目名称不可为空"))
	}
	if len(proj.GitUrl) == 0 {
		return errors.WithStack(errors.New("git url不可为空"))
	}

	if checkRepeat(ctx, proj, creator) {
		return errors.WithStack(fmt.Errorf("项目/%s/%s已存在", creator.Username, proj.ProjectName))
	}

	return nil
}

// 检查项目名称是否重复
func checkRepeat(ctx context.Context, proj *project_idl.Project, creator *dao.User) bool {
	projDao, _ := FindProject(ctx, creator.Username, proj.ProjectName)
	return projDao != nil
}
