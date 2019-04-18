package store

import (
//	"context"

	"github.com/DouwaIO/hairtail/src/model"
)

type Store interface {
	// GetData gets a user by unique ID.
	GetData(string, string) (*model.Data, error)
	// CreateData creates a new user account.
	CreateData(*model.Data) error

	// UpdateData updates a user account.
	UpdateData(*model.Data) error

	// DeleteData deletes a user account.
	DeleteData(*model.Data) error

	GetPipeline(string) (*model.Pipeline, error)
	// CreateData creates a new user account.
	CreatePipeline(*model.Pipeline) error

	// UpdateData updates a user account.
	UpdatePipeline(*model.Pipeline) error

	// DeleteData deletes a user account.
	DeletePipeline(*model.Pipeline) error

	GetSchema(string) (*model.Schema, error)
	// CreateData creates a new user account.
	CreateSchema(*model.Schema) error

	// UpdateData updates a user account.
	UpdateSchema(*model.Schema) error

	// DeleteData deletes a user account.
	DeleteSchema(*model.Schema) error

	GetService(string, string) (*model.Service, error)

	GetServiceList(string) ([]*model.Service, error)
	// CreateService creates a new user account.
	CreateService(*model.Service) error

	// UpdateService updates a user account.
	UpdateService(*model.Service) error

	// DeleteService deletes a user account.
	DeleteService(*model.Service) error

}
