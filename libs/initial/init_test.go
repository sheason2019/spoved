package initial_test

import (
	"context"
	"testing"
	"time"

	"github.com/sheason2019/spoved/libs/initial"
)

func TestInitial(t *testing.T) {
	initial.Initial(context.TODO())

	time.Sleep(time.Second * 20)
}
