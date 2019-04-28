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

func (db *datastore) LogFind(proc *model.Proc) (string, error) {
	data := new(model.LogData)
	err := db.Where("log_job_id = ?",proc.ID).First(&data).Error
	return data.Data, err
	//buf := bytes.NewBuffer(data.Data)
	//return ioutil.NopCloser(buf), err
}

func (db *datastore) LogSave(proc *model.Proc, message string) error {
	data := new(model.LogData)
	err := db.Where("log_job_id = ?",proc.ID).First(&data).Error
	if err != nil {
		data = &model.LogData{ProcID: proc.ID}
	}
	data.Data = message
	err = db.Save(data).Error
	return err

}
