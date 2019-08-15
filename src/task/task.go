package task

import (
	"errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// "github.com/DouwaIO/hairtail/src/model"
)

type Params struct {
	Settings map[string]interface{}
    Data     []byte
	DB			*gorm.DB
}

type Result struct {
    Data        []byte
    SplitData   map[string][]byte
}

type Plugin struct {
	Type	string
	Settings map[string]interface{}
}

func (p *Plugin) Run(data []byte) (*Result, error) {
	log.Debugf("Task %s running...", p.Type)

	params := Params{
		Data:     data,
		Settings: p.Settings,
	}

    switch p.Type {
    case "even":
        return Even(&params)
    case "select":
		return Select(&params)
    case "accumulate":
		// db, err := gorm.Open("postgres", "host=47.110.154.127 port=30011 user=postgres dbname=postgres sslmode=disable password=huansi@2017")
		db, err := gorm.Open("postgres", "host=47.110.154.127 port=30172 user=postgres dbname=hairtail sslmode=disable password=huansi@2017")
		if err != nil {
			log.Errorf("Pipeline run task error: %s", err)
			return nil, err
		}

		// err = db.AutoMigrate(&model.RemoteData{},).Error
		// if err != nil {
		// 	log.Errorf("db auto migrate error: %s", err)
		// 	return nil, err
		// }

		tx := db.Begin()
		params.DB = tx
		result, err := Accumulate(&params)
		if err != nil {
			tx.Rollback()
			log.Errorf("accumulate error: %s", err)
			return nil, err
		}
		tx.Commit()
		db.Close()

		return result, nil
    }
    return nil, errors.New("plugin not fonded")
}
