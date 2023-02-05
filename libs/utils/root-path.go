package utils

import (
	"path"
	"runtime"

	"github.com/sheason2019/spoved/libs/env"
)

func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)

	if !env.IS_PRODUCT {
		return path.Dir(path.Dir(path.Dir(filename)))
	} else {
		return ""
	}
}
