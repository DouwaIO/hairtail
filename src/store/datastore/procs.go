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
)

func (db *datastore) ProcLoad(id int64) (*model.Proc, error) {
	proc := new(model.Proc)
	err := db.Where("proc_id = ?",id).First(&proc).Error
	return proc, err
}

func (db *datastore) ProcFind(build *model.Build, pid int) (*model.Proc, error) {
	proc := new(model.Proc)
	err := db.Where("proc_build_id = ? AND proc_pid = ?",build.ID, pid).First(&proc).Error
	return proc, err
}

func (db *datastore) ProcChild(build *model.Build, pid int, child string) (*model.Proc, error) {
	proc := new(model.Proc)
	err := db.Where("proc_build_id = ? AND proc_ppid = ? AND proc_name = ?", build.ID, pid, child).First(&proc).Error
	return proc, err
}

func (db *datastore) ProcList(build *model.Build) ([]*model.Proc, error) {
	list := []*model.Proc{}
	err := db.Where("proc_build_id = ?",build.ID).Order("proc_id ASC").Find(&list).Error
	return list, err
}

func (db *datastore) ProcCreate(procs []*model.Proc) error {
	for _, proc := range procs {
		if err := db.Create(proc).Error; err != nil {
			return err
		}
	}
	return nil

}

func (db *datastore) ProcUpdate(proc *model.Proc) error {
	var count int
	err := db.Model(&model.Proc{}).Where("proc_id = ?",proc.ID).Count(&count).Error
	if err != nil || count == 0{
		return nil
	}

	return db.Save(proc).Error
}

func (db *datastore) ProcClear(build *model.Build) (err error) {
	// 开启事务

	tx := db.Begin()
	//err = tx.Where("file_build_id = ?", build.ID).Delete(model.File{}).Error
	//if err != nil{
	//	tx.Rollback()
	//	return err
	//}
	err = tx.Where("proc_build_id = ?", build.ID).Delete(model.Proc{}).Error
	if err != nil{
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()
	return nil
}
