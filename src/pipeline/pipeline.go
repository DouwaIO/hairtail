package pipeline

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// "github.com/DouwaIO/hairtail/src/model"
	// "github.com/DouwaIO/hairtail/src/utils"
	"github.com/DouwaIO/hairtail/src/task"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

type Pipeline struct {
	TargetDB *gorm.DB
	Tasks    []*yaml.Task
}

func (p *Pipeline) Run(data []byte) error {
	log.Info("Pipeline running...")

	for _, t := range p.Tasks {
		tk := task.Plugin{
			TargetDB: p.TargetDB,
			Type:     t.Type,
			Settings: t.Settings,
		}
		result, err := tk.Run(data)
		if err != nil {
			log.Errorf("Pipeline run task error: %s", err)
			log.Info("Pipeline end")
			return err
		}

		if result != nil && result.Data != nil {
			data = result.Data
		}
	}

	log.Info("Pipeline end")
	return nil
}
