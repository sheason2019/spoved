package project_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/dbc"
	project_idl "github.com/sheason2019/spoved/libs/idl-lib/project"
)

func CreateProject(ctx context.Context, proj *project_idl.Project, creator *dao.User) (*dao.Project, error) {
	// 校验逻辑
	err := validate(ctx, proj, creator)
	if err != nil {
		return nil, err
	}

	client := dbc.DB
	projDao := &dao.Project{
		ProjectName: proj.ProjectName,
		Describe:    proj.Describe,
		GitUrl:      proj.GitUrl,
		Creator:     *creator,
	}

	err = client.WithContext(ctx).Save(projDao).Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return projDao, nil
}
