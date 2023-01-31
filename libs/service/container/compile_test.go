package container_service_test

import (
	"testing"

	container_service "github.com/sheason2019/spoved/libs/service/container"
)

func TestCompile(t *testing.T) {
	e := container_service.Compile("node:16-alpine", "Patch", "master", 1)
	if e != nil {
		e.Panic()
	}
}
