package store

import (
	"context"

	"github.com/DouwaIO/hairtail/src/model"
)

type Store interface {
	// GetUser gets a user by unique ID.
	GetUser(int64) (*model.User, error)
	// CreateUser creates a new user account.
	CreateUser(*model.User) error

	// UpdateUser updates a user account.
	UpdateUser(*model.User) error

	// DeleteUser deletes a user account.
	DeleteUser(*model.User) error

}

// GetUser gets a user by unique ID.
func GetUser(c context.Context, id int64) (*model.User, error) {
	return FromContext(c).GetUser(id)
}
