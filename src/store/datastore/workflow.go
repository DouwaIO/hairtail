package datastore

import (
	"github.com/DouwaIO/hairtail/src/model"
)

func (db *datastore) GetWorkflows(search string) ([]*model.Workflow, error) {
	workflows := []*model.Workflow{}
	err := db.Where("name like ?", "%"+search+"%").Order("name").Find(&workflows).Error
	return workflows, err
}

func (db *datastore) GetWorkflow(id string) (*model.Workflow, error) {
	var workflow = new(model.Workflow)
	err := db.Where("id = ?", id).First(&workflow).Error
	return workflow, err
}

func (db *datastore) CreateWorkflow(workflow *model.Workflow) error {
	err := db.Create(workflow).Error
	return err
}

func (db *datastore) UpdateWorkflow(workflow *model.Workflow) error {
	var count int
	err := db.Where("id = ?", workflow.ID).Find(&model.Workflow{}).Count(&count).Error
	if err != nil || count == 0 {
		return nil
	}

	return db.Save(workflow).Error
}

func (db *datastore) DeleteWorkflow(id string) error {
	err := db.Where("id = ?", id).Delete(&model.Workflow{}).Error
	return err
}
