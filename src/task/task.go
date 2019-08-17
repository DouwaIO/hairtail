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
	DB       *gorm.DB
}

type Result struct {
	Data      []byte
	SplitData map[string][]byte
}

type Plugin struct {
	TargetDB *gorm.DB
	Type     string
	Settings map[string]interface{}
}

func (p *Plugin) Run(data []byte) (*Result, error) {
	log.WithFields(log.Fields{"type": p.Type}).Info("task running...")

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
		tx := p.TargetDB.Begin()
		defer func() {
			if err := recover(); err != nil {
				tx.Rollback()
			}
		}()
		if err := tx.Error; err != nil {
			return nil, err
		}
		// if err := tx.Exec(`set transaction isolation level repeatable read`).Error; err != nil {
		if err := tx.Exec(`set transaction isolation level read committed`).Error; err != nil {
			return nil, err
		}

		params.DB = tx
		result, err := Accumulate(&params)
		if err != nil {
			return nil, err
		}
		tx.Commit()

		return result, nil
	}
	return nil, errors.New("plugin not fonded")
}
