package pipeline

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	// "github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/pipeline/queue"
	// "github.com/DouwaIO/hairtail/src/utils"
	"github.com/DouwaIO/hairtail/src/task"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

type Pipeline struct {
	TargetDB *gorm.DB
	Tasks    []*yaml.Task
}

var Queue queue.Queue

func (p *Pipeline) Run(data []byte) error {
	log.Info("Pipeline running...")

	// ctx := context.Background()
	// gen_id := utils.GeneratorId()
	// task2 := &queue.Task{
	// 	ID:   gen_id,
	// 	Data: data,
	// }
	// Queue.Push(ctx, task2)
	// fn, err := queue.CreateFilterFunc("filter")
	// if err != nil {
	// 	fmt.Println("s", err)
	// }
	// Queue.Poll(ctx, fn, gen_id)

	for _, t := range p.Tasks {
		tk := task.Plugin{
			TargetDB: p.TargetDB,
			Type:     t.Type,
			Settings: t.Settings,
		}
		result, err := tk.Run(data)
		log.Debugf("Task end")
		if err != nil {
			log.Errorf("Pipeline run task error: %s", err)
			return err
		}
		// log.Debugf("Pipeline task result: %s", result)

		if result != nil && result.Data != nil {
			data = result.Data
		}
	}
	// Queue.Done(ctx, gen_id)
	// log.Debugf("done: %s\n", Queue.Info(ctx))
	return nil
}
