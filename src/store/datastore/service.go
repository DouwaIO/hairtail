// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"github.com/DouwaIO/hairtail/src/model"
	// "gitlab.com/douwa/registry/dougo/src/dougo/store/datastore/sql"
	// "github.com/russross/meddler"
	// "errors"
)

func (db *datastore) GetService(name string, pipeline string) (*model.Service, error) {
	var usr = new(model.Service)
	err := db.Where("name = ? and pipeline_id = ?", name, pipeline).First(&usr).Error
	return usr,err
}

func (db *datastore) GetServiceAllList() ([]*model.Service, error) {

	data := []*model.Service{}
	err := db.Where("type = ?", "MQ").Find(&data).Error
	return data,err
}

func (db *datastore) GetServiceList(pipeline string) ([]*model.Service, error) {

	data := []*model.Service{}
	err := db.Where("pipeline_id = ?", pipeline).Find(&data).Error
	return data,err
}

func (db *datastore) CreateService(data *model.Service) error {
	err := db.Create(data).Error
	return err
}

func (db *datastore) UpdateService(data *model.Service) error {
	var count int
	err := db.Model(&model.Service{}).Where("service_id = ?",data.ID).Count(&count).Error
	if err != nil || count == 0{
		return nil
	}

	return db.Save(data).Error
}

func (db *datastore) DeleteService(data *model.Service) error {
	err := db.Where("service_id = ?",data.ID).Delete(model.Service{}).Error
	return err
}
