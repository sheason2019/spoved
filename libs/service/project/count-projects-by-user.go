package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CountProjectsByUser(usr *ent.User) (int, error) {
	client := dbc.GetClient()

	count, err := client.Project.Query().
		Where(
			project.HasCreatorWith(
				user.IDEQ(usr.ID),
			),
		).
		Count(context.Background())
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}
