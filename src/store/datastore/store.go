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
	// "os"

	"github.com/DouwaIO/hairtail/src/store"

	"github.com/DouwaIO/hairtail/src/model"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// datastore is an implementation of a model.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*gorm.DB

	driver string
	config string
}

// var db *gorm.DB

// New creates a database connection for the given driver and datasource
// and returns a new Store.
func New(driver string) store.Store {
	return &datastore{
		DB:     open(driver),
		driver: driver,
		// config: config,
	}
}

// From returns a Store using an existing database connection.
// func From(db *sql.DB) store.Store {
// 	return &datastore{DB: db}
// }

// open opens a new database connection with the specified
// driver and connection string and returns a store.
func open(driver string) *gorm.DB {
	//db,err := gorm.Open("sqlite3", "test.db")
	// os.Getenv("DRONE_DATABASE_DATASOURCE")
	DRONE_DATABASE_DATASOURCE := driver
	db,err := gorm.Open("postgres", DRONE_DATABASE_DATASOURCE)
		if err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("database ping attempts failed")
		}
		logrus.Infof("连接成功了")
		if err := setupDatabase(driver, db); err != nil {
			logrus.Errorln(err)
			logrus.Fatalln("migration failed")
		}
	return db
}


// helper function to setup the databsae by performing
// automated database migration steps.
func setupDatabase(driver string, db *gorm.DB) error {
	db.Set("gorm:table_options", "charset=utf8")
	return db.AutoMigrate(&model.Schema{},
			      &model.Service{},
			      &model.Pipeline{},
			      &model.Data{}).Error
}

