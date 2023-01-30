package file_service_test

import (
	"fmt"
	"testing"

	file_service "github.com/sheason2019/spoved/libs/service/file"
)

func TestGitClone(t *testing.T) {
	output, err := file_service.GitClone("https://github.com/sheason2019/node-template", "/repos/node-template", "master")
	if err != nil {
		err.Panic()
	}
	fmt.Println(output)
}
