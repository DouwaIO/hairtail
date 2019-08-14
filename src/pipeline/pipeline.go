package pipeline

import (
	"context"
	"fmt"
	"log"

	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/pipeline/queue"
	"github.com/DouwaIO/hairtail/src/utils"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
	"github.com/DouwaIO/hairtail/src/task/plugins/even"
	"github.com/DouwaIO/hairtail/src/task/plugins/selec"
	"github.com/DouwaIO/hairtail/src/task/plugins/accumulate"
)

type Pipeline struct {
	Tasks []*yaml.Task
}

var Queue queue.Queue

func (p *Pipeline) Run(data []byte) (string, error) {
	//parsed, err := yaml.ParseString(q.config)
	//if err != nil {
	//	return errors.New("yaml type error")
	//}
	log.Printf("hello world")
	ctx := context.Background()
	gen_id := utils.GeneratorId()
	task2 := &queue.Task{
		ID:   gen_id,
		Data: data,
	}
	Queue.Push(ctx, task2)
	fn, err := queue.CreateFilterFunc("filter")
	if err != nil {
		fmt.Println("s", err)
	}
	Queue.Poll(ctx, fn, gen_id)
	// log.Printf("poll: %s\n", Queue.Info(ctx))

	if len(p.Tasks) > 0 {
		for _, t := range p.Tasks {
            tk := GetTaskPlugin(t)
            params := task.Params{
                Settings: t.Settings,
                Data:     data,
            }
			result, err := tk.Run(&params)
			if err != nil {
				fmt.Printf("task run error: %s\n", err)
			}
			data = result.Data
		}
	}
	Queue.Done(ctx, gen_id)
	log.Printf("done: %s\n", Queue.Info(ctx))
	return model.StatusSuccess, nil
}

func GetTaskPlugin(t *yaml.Task) task.Plugin {
    var tk task.Plugin
    switch t.Type {
    case "even":
        tk = new(even.Plugin)
    case "select":
        tk = new(selec.Plugin)
    case "accumulate":
        tk = new(accumulate.Plugin)
    }
    return tk
}
