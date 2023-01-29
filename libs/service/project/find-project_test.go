package project_service_test

import (
	"fmt"
	"testing"

	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func TestFindProjectTest(t *testing.T) {
	username := "sheason"
	projName := "node-template"

	proj, e := project_service.FindProject(username, projName)
	if e != nil {
		e.Panic()
	}

	fmt.Println(proj)
}
