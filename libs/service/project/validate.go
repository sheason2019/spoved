package project_service

import (
	"errors"

	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
)

func validate(proj *project.Project) *exception.Exception {
	if len(proj.ProjectName) == 0 {
		return exception.New(errors.New("项目名称不可为空"))
	}
	if len(proj.GitUrl) == 0 {
		return exception.New(errors.New("git url不可为空"))
	}

	return nil
}
