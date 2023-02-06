package utils

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func TimeoutFunc(ctx context.Context, fn func(ctx context.Context, cancel func()), duration time.Duration) error {
	toCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	go fn(toCtx, cancel)

	select {
	case <-toCtx.Done():
		break
	case <-time.After(duration * time.Millisecond):
		return errors.WithStack(errors.New("拉取仓库超时"))
	}

	return nil
}
