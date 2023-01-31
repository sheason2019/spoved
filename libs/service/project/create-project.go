package project_service

import (
	"context"
	"fmt"
	"time"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
	project_idl "github.com/sheason2019/spoved/libs/idl-lib/project"
)

func CreateProject(proj *project_idl.Project, creator *ent.User) (*ent.Project, *exception.Exception) {
	// 校验逻辑
	e := validate(proj, creator)
	if e != nil {
		return nil, e.Wrap()
	}

	client := dbc.GetClient()

	entProj, err := client.Project.
		Create().
		SetProjectName(proj.ProjectName).
		SetDescribe(proj.Describe).
		SetGitURL(proj.GitUrl).
		SetDirPath(fmt.Sprintf("repos/%s/%s", creator.Username, proj.ProjectName)).
		SetCreatedAt(time.Now()).
		AddCreator(creator).
		Save(context.Background())

	if err != nil {
		return nil, exception.New(err)
	}

	return entProj, nil
}
