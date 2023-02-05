package project_service_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"

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

func TestErrorStack(t *testing.T) {
	err := createErr()
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func createErr() error {
	return errors.WithStack(errors.Errorf("Test Error"))
}
