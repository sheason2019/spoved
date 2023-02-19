package main

import (
	"fmt"

	"github.com/sheason2019/spoved/libs/router"
)

func main() {
	fmt.Println("START SPOVED")
	r := router.SetupRouter()

	r.Run(":80")
}
