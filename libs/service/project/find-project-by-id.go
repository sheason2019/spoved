package project_service

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/project"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindProjectById(id int) (*ent.Project, *exception.Exception) {
	client := dbc.GetClient()

	proj, err := client.Project.Query().Where(project.IDEQ(id)).First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, exception.New(err)
	}

	return proj, nil
}
