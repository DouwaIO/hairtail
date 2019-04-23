package store

import (
//	"context"

	"github.com/DouwaIO/hairtail/src/model"
)

type Store interface {
	GetData(string, string) (*model.Data, error)

	CreateData(*model.Data) error

	UpdateData(*model.Data) error

	DeleteData(*model.Data) error

	GetPipeline(string) (*model.Pipeline, error)

	CreatePipeline(*model.Pipeline) error

	UpdatePipeline(*model.Pipeline) error

	DeletePipeline(*model.Pipeline) error

	GetSchema(string) (*model.Schema, error)

	CreateSchema(*model.Schema) error

	UpdateSchema(*model.Schema) error

	DeleteSchema(*model.Schema) error

	GetService(string, string) (*model.Service, error)

	GetServiceAllList() ([]*model.Service, error)

	GetServiceList(string) ([]*model.Service, error)

	CreateService(*model.Service) error

	UpdateService(*model.Service) error

	DeleteService(*model.Service) error

	GetStep(string, int64) (*model.Step, error)

	GetStepList(string) ([]*model.Step, error)
	CreateStep(*model.Step) error

	UpdateStep(*model.Step) error

	DeleteStep(*model.Step) error

	GetBuildList(string) ([]*model.Build, error)
	CreateBuild(*model.Build) error

	UpdateBuild(*model.Build) error

	DeleteBuild(*model.Build) error

}
