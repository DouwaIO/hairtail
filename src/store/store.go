package store

import (
	"github.com/DouwaIO/hairtail/src/model"
)

type Store interface {
	GetSchema(string) (*model.Schema, error)
	CreateSchema(*model.Schema) error
	UpdateSchema(*model.Schema) error
	DeleteSchema(*model.Schema) error

	GetStep(string, int64) (*model.Step, error)
	GetStepList(string) ([]*model.Step, error)
	CreateStep(*model.Step) error
	UpdateStep(*model.Step) error
	DeleteStep(*model.Step) error

	GetBuildList(string) ([]*model.Build, error)
	CreateBuild(*model.Build) error
	UpdateBuild(*model.Build) error
	DeleteBuild(*model.Build) error

    GetPipelines(string) ([]*model.Pipeline, error)
    GetPipeline(string) (*model.Pipeline, error)
    CreatePipeline(*model.Pipeline) error
    UpdatePipeline(*model.Pipeline) error
    DeletePipeline(string) error
}
