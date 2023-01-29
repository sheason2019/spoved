package project_service

import (
	"context"
	"errors"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
	project_idl "github.com/sheason2019/spoved/libs/idl-lib/project"
)

func validate(proj *project_idl.Project, creator *ent.User) *exception.Exception {
	if len(proj.ProjectName) == 0 {
		return exception.New(errors.New("项目名称不可为空"))
	}
	if len(proj.GitUrl) == 0 {
		return exception.New(errors.New("git url不可为空"))
	}

	e := checkRepeat(proj, creator)
	if e != nil {
		return e.Wrap()
	}

	return nil
}

// 检查项目名称是否重复
func checkRepeat(proj *project_idl.Project, creator *ent.User) *exception.Exception {
	client := dbc.GetClient()
	count, err := client.Project.Query().
		Where(project.ProjectNameEQ(proj.ProjectName)).
		Where(project.HasCreatorWith(user.IDEQ(creator.ID))).
		Count(context.Background())

	if err != nil {
		return exception.New(err)
	}

	if count > 0 {
		return exception.New(errors.New("项目名称重复"))
	}

	return nil
}
