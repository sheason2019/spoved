package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	project_idl "github.com/sheason2019/spoved/libs/idl-lib/project"
)

func validate(ctx context.Context, proj *project_idl.Project, creator *dao.User) error {
	if len(proj.ProjectName) == 0 {
		return errors.WithStack(errors.New("项目名称不可为空"))
	}
	if len(proj.GitUrl) == 0 {
		return errors.WithStack(errors.New("git url不可为空"))
	}

	return checkRepeat(ctx, proj, creator)
}

// 检查项目名称是否重复
func checkRepeat(ctx context.Context, proj *project_idl.Project, creator *dao.User) error {
	client := dbc.GetClient()
	var count int64
	err := client.WithContext(ctx).
		Model(&dao.Project{}).
		Where("project_name = ?", proj.ProjectName).
		Where("creator_id = ?", creator.ID).
		Count(&count).
		Error

	if err != nil {
		return errors.WithStack(err)
	}

	if count > 0 {
		return errors.WithStack(errors.New("项目名称重复"))
	}

	return nil
}
