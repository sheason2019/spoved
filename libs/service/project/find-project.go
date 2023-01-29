package project_service

import (
	"context"
	"errors"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindProject(username, projName string) (*ent.Project, *exception.Exception) {
	client := dbc.GetClient()

	proj, err := client.Project.Query().
		Where(
			project.ProjectNameEQ(projName),
			project.HasCreatorWith(
				user.UsernameEQ(username),
			),
		).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, exception.New(errors.New("未找到指定的项目" + `/` + username + `/` + projName))
	}
	if err != nil {
		return nil, exception.New(err)
	}

	return proj, nil
}
