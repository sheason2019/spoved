package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindProjectById(id int) (*ent.Project, error) {
	client := dbc.GetClient()

	proj, err := client.Project.Query().Where(project.IDEQ(id)).First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return proj, nil
}
