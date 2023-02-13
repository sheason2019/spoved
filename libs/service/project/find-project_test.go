package project_service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"

	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func TestFindProjectTest(t *testing.T) {
	username := "sheason"
	projName := "node-template"

	proj, err := project_service.FindProject(context.TODO(), username, projName)
	if err != nil {
		t.Errorf("%+v", err)
		return
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
