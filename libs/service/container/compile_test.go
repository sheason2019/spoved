package container_service_test

import (
	"fmt"
	"testing"

	container_service "github.com/sheason2019/spoved/libs/service/container"
)

func TestCompile(t *testing.T) {
	output, e := container_service.Compile("node:16-alpine", "Patch", "master", 1, "sheason")
	fmt.Println("output::\n", output)
	if e != nil {
		t.Errorf("%+v", e)
		return
	}
}

func TestCompileRun(t *testing.T) {
	output, e := container_service.CompileRun("/repos/sheason/node-template/0.0.1")
	fmt.Println(output)
	if e != nil {
		t.Errorf("%+v", e)
		return
	}
}
