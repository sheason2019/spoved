package main

import (
	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	r := router.SetupProxy()

	r.Run(":80")
}
