package task

import (
	"context"
)

type Task interface {
	User(token string, t int) (*model.User, error)
}

func User(c context.Context, token string, t int) (*model.User, error) {
	return FromContext(c).User(token, t)
}
