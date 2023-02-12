package account_service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindUserByUsername(username string) (*ent.User, error) {
	client := dbc.GetClient()

	user, err := client.User.
		Query().
		Where(user.UsernameEQ(username)).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
