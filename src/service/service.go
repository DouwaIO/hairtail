package service

import (
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/store"
	"github.com/DouwaIO/hairtail/src/pipeline"
)

type Service struct {
	Name          string
	Desc          string
	Type          string
	Settings      map[string]interface{}
    Steps         []*yaml.Task
    // database store
    Store         *store.Store
}


func (s *Service) Run() error {
	if s.Type == "MQ" {
		go MQ(s)
	} else if s.Type == "DB" {
		// go DB(service.Settings["db_type"].(string), service.Settings["host"].(string), service.Settings["port"].(string), service.Settings["user"].(string), service.Settings["pwd"].(string), service.Settings["name"].(string), service.Settings["table"].(string), service.Settings["column"].(string), 1, 1)
	}

	return nil
}

func (s *Service) RunStep(data []byte) error {
	// newdata := &model.Build{
	// 	Service: service,
	// 	Data:    string(d.Body),
	// 	//Status: model.StatusPending,
	// 	Status:     model.StatusRunning,
	// 	Timestamp:  currentTime,
	// 	Timestamp2: int64(0),
	// }
	// err = v.CreateBuild(newdata)
	// if err != nil {
	// 	log.Printf("add data error")
	// }
	// status := yaml.Pipeline(data, d.Body)
	// currentTime = time.Now().Unix()
	// newdata.Status = status
	// newdata.Timestamp2 = currentTime
	// err = v.UpdateBuild(newdata)
	// if err != nil {
	// 	log.Printf("add data error")
	// }

    p := pipeline.Pipeline{
        Tasks:  s.Steps,
    }

    _, err := p.Run(data)
    return err
}
