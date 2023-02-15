package main

import (
	"context"

	"github.com/sheason2019/spoved/libs/initial"
	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	ctx := context.TODO()
	initial.Initial(ctx)

	r := router.SetupRouter()

	r.Run(":80")
}
