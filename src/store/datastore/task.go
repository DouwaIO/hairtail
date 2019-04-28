package datastore

import (
	"github.com/DouwaIO/hairtail/src/model"
)

func (db *datastore) TaskList() ([]*model.Task, error) {
	data := []*model.Task{}
	err := db.Find(&data).Error
	return data, err
}

func (db *datastore) TaskInsert(task *model.Task) error {
	err := db.Create(task).Error
	return err
}

func (db *datastore) TaskDelete(id string) error {
	err := db.Where("task_id = ?",id).Delete(model.Task{}).Error
	return err
}
