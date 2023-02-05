package file_service_test

import (
	"fmt"
	"testing"

	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func TestGitClone(t *testing.T) {
	output, err := file_service.GitClone("https://github.com/sheason2019/node-template", "/repos/node-template/0.0.1", "master")
	if err != nil {
		t.Errorf("%+v", err)
	}
	fmt.Println(output)
}
