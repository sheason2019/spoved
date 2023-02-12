package project_controller_test

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sheason2019/spoved/libs/idl-lib/project"
	test_utils "github.com/sheason2019/spoved/libs/test-utils"
)

func TestGetProject(t *testing.T) {
	payload := project.GetProjectPayload{
		Username:    "sheason",
		ProjectName: "node-template",
	}

	proj, w, e := test_utils.HttpTestWithRecorder[project.Project]("GET", project.ProjectApiDefinition.GET_PROJECT_PATH, payload)

	assert.Equal(t, w.Code, 200)

	if e != nil {
		t.Error(e)
	}

	fmt.Println(proj)
}
