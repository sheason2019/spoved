package compile_service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sheason2019/spoved/libs/env"
	compile_service "github.com/sheason2019/spoved/libs/service/compile"
	project_service "github.com/sheason2019/spoved/libs/service/project"
)

func TestCompile(t *testing.T) {
	proj, err := project_service.FindProjectById(context.TODO(), 1)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}

	output, e := compile_service.CompileRun("node:16-alpine", "Patch", "0.0.2", proj, "sheason")
	fmt.Println("output::\n", output)
	if e != nil {
		t.Errorf("%+v", e)
		return
	}
}

func TestCompileRun(t *testing.T) {
	output, e := compile_service.CompileRunBuild(context.Background(), env.DataRoot+"/repos/sheason/node-template/0.0.2")
	fmt.Println(output)
	if e != nil {
		t.Errorf("%+v", e)
		return
	}
}
