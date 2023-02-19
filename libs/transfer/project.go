package transfer

import (
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

func ProjectToIdl(proj *dao.Project) *project.Project {
	return &project.Project{
		Id:          int(proj.ID),
		ProjectName: proj.ProjectName,
		GitUrl:      proj.GitUrl,
		Describe:    proj.Describe,
		Owner:       proj.Creator.Username,
	}
}

func ProjectsToIdl(projs []*dao.Project) []*project.Project {
	targets := make([]*project.Project, len(projs))
	for i, v := range projs {
		targets[i] = ProjectToIdl(v)
	}

	return targets
}
