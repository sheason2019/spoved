package transfer

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

func ProjectToIdl(proj *ent.Project) project.Project {
	target := project.Project{
		Id:          proj.ID,
		ProjectName: proj.ProjectName,
		GitUrl:      proj.GitURL,
		Describe:    proj.Describe,
	}

	creator := proj.QueryCreator().FirstX(context.Background())
	target.Owner = creator.Username

	return target
}

func ProjectsToIdl(projs []*ent.Project) []project.Project {
	targets := make([]project.Project, len(projs))
	for i, v := range projs {
		targets[i] = ProjectToIdl(v)
	}

	return targets
}
