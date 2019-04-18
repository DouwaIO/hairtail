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

func (db *datastore) GetSchema(name string,) (*model.Schema, error) {
	var usr = new(model.Schema)
	err := db.Where("name = ?", name).First(&usr).Error
	return usr,err
}

func (db *datastore) CreateSchema(data *model.Schema) error {
	err := db.Create(data).Error
	return err
}

func (db *datastore) UpdateSchema(data *model.Schema) error {
	var count int
	err := db.Model(&model.Schema{}).Where("schema_id = ?",data.ID).Count(&count).Error
	if err != nil || count == 0{
		return nil
	}

	return db.Save(data).Error
}

func (db *datastore) DeleteSchema(data *model.Schema) error {
	err := db.Where("schema_id = ?",data.ID).Delete(model.Schema{}).Error
	return err
}
