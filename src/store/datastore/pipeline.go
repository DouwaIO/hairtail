package datastore

import (
	"github.com/DouwaIO/hairtail/src/model"
)

func (db *datastore) GetPipelines(search string) ([]*model.Pipeline, error) {
	pipelines := []*model.Pipeline{}
	err := db.Where("name like ?", "%"+search+"%").Order("name").Find(&pipelines).Error
	return pipelines, err
}

func (db *datastore) GetPipeline(id string) (*model.Pipeline, error) {
	var pipeline = new(model.Pipeline)
	err := db.Where("id = ?", id).First(&pipeline).Error
	return pipeline, err
}

func (db *datastore) CreatePipeline(pipeline *model.Pipeline) error {
	err := db.Create(pipeline).Error
	return err
}

func (db *datastore) UpdatePipeline(pipeline *model.Pipeline) error {
	var count int
	err := db.Where("id = ?", pipeline.ID).Find(&model.Pipeline{}).Count(&count).Error
	if err != nil || count == 0 {
		return nil
	}

	return db.Save(pipeline).Error
}

func (db *datastore) DeletePipeline(id string) error {
	err := db.Where("id = ?", id).Delete(&model.Pipeline{}).Error
	return err
}
