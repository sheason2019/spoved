package regist_service

import (
	"context"
	"time"

	"github.com/sheason2019/spoved/ent"
	"github.com/sheason2019/spoved/exceptions/exception"
	"github.com/sheason2019/spoved/libs/dbc"
)

func CreateUser(username, password, salt string) (*ent.User, *exception.Exception) {
	client := dbc.GetClient()
	usr, err := client.User.Create().
		SetUsername(username).
		SetPasswordHash(password).
		SetPasswordSalt(salt).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, exception.New(err)
	}

	return usr, nil
}
