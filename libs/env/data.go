package env

import (
	"path"
	"runtime"
)

var DataRoot string

func init() {
	_, filename, _, _ := runtime.Caller(0)

	if !IS_PRODUCT {
		DataRoot = path.Dir(path.Dir(path.Dir(filename))) + "/data"
	} else {
		DataRoot = "/data"
	}
}
