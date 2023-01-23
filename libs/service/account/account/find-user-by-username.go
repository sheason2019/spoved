package account_service

import (
	"context"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/ent/user"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func FindUserByUsername(username string) (*ent.User, *exception.Exception) {
	client := dbc.GetClient()

	users, err := client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Limit(1).
		All(context.Background())
	if err != nil {
		return nil, exception.New(err)
	}
	if len(users) > 0 {
		return users[0], nil
	}

	return nil, nil
}
