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

func (db *datastore) GetUser(id int64) (*model.User, error) {
	// var usr = new(model.User)
	// var err = meddler.Load(db, "users", usr, id)
	// return usr, err

	var usr = new(model.User)
	err := db.Where("user_id = ?",id).First(&usr).Error
	return usr,err
}

func (db *datastore) CreateUser(user *model.User) error {
	// return meddler.Insert(db, "users", user)
	err := db.Create(user).Error
	return err
}

func (db *datastore) UpdateUser(user *model.User) error {
	// return meddler.Update(db, "users", user)

	// 由于.save 在不存在的时候，会创建新的数据，所以需要判断是否存在
	var count int
	err := db.Model(&model.User{}).Where("user_id = ?",user.ID).Count(&count).Error
	if err != nil || count == 0{
		return nil
	}

	return db.Save(user).Error
}

func (db *datastore) DeleteUser(user *model.User) error {
	// stmt := sql.Lookup(db.driver, "user-delete")
	// _, err := db.Exec(stmt, user.ID)
	// return err

	// 必须确保主键有值，否则会将数据全部删除（大坑啊）
	// if user.ID == 0{
	// 	return errors.New("要删除的对象不存在，请传递ID")
	// }
	// err := db.Delete(user).Error
	// return err
	err := db.Where("user_id = ?",user.ID).Delete(model.User{}).Error
	return err
}
