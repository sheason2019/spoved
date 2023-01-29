package project_service

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CountProjectsByUser(usr *ent.User) (int, *exception.Exception) {
	client := dbc.GetClient()

	count, err := client.Project.Query().
		Where(
			project.HasCreatorWith(
				user.IDEQ(usr.ID),
			),
		).
		Count(context.Background())
	if err != nil {
		return 0, exception.New(err)
	}

	return count, nil
}
