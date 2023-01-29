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

	user, err := client.User.
		Query().
		Where(user.UsernameEQ(username)).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, exception.New(err)
	}

	return user, nil
}
