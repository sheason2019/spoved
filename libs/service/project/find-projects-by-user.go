package project_service

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
	"github.com/sheason2019/spoved/libs/idl-lib/common"
)

// 根据用户名搜索项目
func FindProjectsByUser(usr *ent.User, pagination *common.Pagination) ([]*ent.Project, *exception.Exception) {
	client := dbc.GetClient()
	projs, err := client.Project.Query().
		Where(
			project.HasCreatorWith(
				user.IDEQ(usr.ID),
			),
		).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit(pagination.PageSize).
		All(context.Background())
	if err != nil {
		return nil, exception.New(err)
	}

	return projs, nil
}
