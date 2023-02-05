package regist_service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateUser(username, password, salt string) (*ent.User, error) {
	client := dbc.GetClient()
	usr, err := client.User.Create().
		SetUsername(username).
		SetPasswordHash(password).
		SetPasswordSalt(salt).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return usr, nil
}
