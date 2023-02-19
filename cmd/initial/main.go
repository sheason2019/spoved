package main

import (
	"context"
	"fmt"

	"github.com/sheason2019/spoved/libs/initial"
)

func main() {
	fmt.Println("START INITIAL")

	ctx := context.TODO()
	initial.Initial(ctx)
}
