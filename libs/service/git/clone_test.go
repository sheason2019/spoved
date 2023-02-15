package git_service_test

import (
	"testing"

	git_service "github.com/sheason2019/spoved/libs/service/git"
)

func TestGitClone(t *testing.T) {
	err := git_service.GitClone("git@github.com:sheason2019/node-template.git", "/repos/node-template/0.0.1", "master", "sheason")
	if err != nil {
		t.Errorf("%+v", err)
	}
}
