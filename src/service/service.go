package service

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/DouwaIO/hairtail/src/pipeline"
	"github.com/DouwaIO/hairtail/src/store"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

type Service struct {
	Name     string
	Desc     string
	Type     string
	Settings map[string]interface{}
	Steps    []*yaml.Task
	// database store
	Store    *store.Store
	// target database
	TargetDB	*gorm.DB
}

func (s *Service) Run() error {
    log.Debugf("Service Run for: %s", s.Type)
	if s.Type == "mq" {
		go MQ(s)
	} else if s.Type == "db" {
		// go DB(service.Settings["db_type"].(string), service.Settings["host"].(string), service.Settings["port"].(string), service.Settings["user"].(string), service.Settings["pwd"].(string), service.Settings["name"].(string), service.Settings["table"].(string), service.Settings["column"].(string), 1, 1)
	}

	return nil
}

func (s *Service) RunStep(data []byte) error {
	log.Debugf("Service run step")

	// currentTime = time.Now().Unix()
	// newdata.Status = status
	// newdata.Timestamp2 = currentTime

	go func() {
		p := pipeline.Pipeline{
			TargetDB: s.TargetDB,
			Tasks: s.Steps,
		}

		_ = p.Run(data)
	}()

	return nil
}
