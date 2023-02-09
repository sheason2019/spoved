package main

import (
	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	r := router.SetupRouter()

	r.Run(":80")
}
