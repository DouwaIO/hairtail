package datastore

import (
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DouwaIO/hairtail/src/store"
	"github.com/DouwaIO/hairtail/src/model"
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
	//DRONE_DATABASE_DATASOURCE := driver
	db,err := gorm.Open("postgres", driver)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database ping attempts failed")
	}
	log.Infof("database connected")
	if err := setupDatabase(db); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("migration failed")
	}
	return db
}


// helper function to setup the databsae by performing
// automated database migration steps.
func setupDatabase(db *gorm.DB) error {
	db.Set("gorm:table_options", "charset=utf8")
	return db.AutoMigrate(&model.Schema{},
			&model.Pipeline{},
			&model.Build{},
			&model.LogData{},
			&model.Proc{}).Error
}
