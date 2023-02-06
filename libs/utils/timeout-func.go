package utils

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func TimeoutFunc(ctx context.Context, fn func(ctx context.Context), duration time.Duration) error {
	toCtx, cancel := context.WithTimeout(ctx, duration*time.Millisecond)
	defer cancel()

	go fn(toCtx)

	select {
	case <-toCtx.Done():
		break
	case <-time.After(duration * time.Millisecond):
		return errors.WithStack(errors.New("拉取仓库超时"))
	}

	return nil
}
