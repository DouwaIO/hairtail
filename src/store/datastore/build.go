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

//func (db *datastore) GetBuild(build string, data string) (*model.Build, error) {
//	var usr = new(model.Build)
//	err := db.Where("pipeline_id = ? and data_id = ?", build, data).First(&usr).Error
//	return usr,err
//}

func (db *datastore) GetBuildList(service string) ([]*model.Build, error) {
	data := []*model.Build{}
	err := db.Where("service_id = ?", service).Find(&data).Error
	return data,err
}

func (db *datastore) CreateBuild(data *model.Build) error {
	err := db.Create(data).Error
	return err
}

func (db *datastore) UpdateBuild(data *model.Build) error {
	var count int
	err := db.Model(&model.Build{}).Where("build_id = ?",data.ID).Count(&count).Error
	if err != nil || count == 0{
		return nil
	}

	return db.Save(data).Error
}

func (db *datastore) DeleteBuild(data *model.Build) error {
	err := db.Where("build_id = ?",data.ID).Delete(model.Build{}).Error
	return err
}
