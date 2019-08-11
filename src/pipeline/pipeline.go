package pipeline

import (
	"context"
	"fmt"
	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/pipeline/queue"
	"github.com/DouwaIO/hairtail/src/task"
	"github.com/DouwaIO/hairtail/src/utils"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"log"
)

type Pipeline struct {
    Tasks         []*yaml.Task
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
			if _, ok := task.Funcs[t.Type]; ok {
				data2, status := task.CallPipeline(t, data)
				if status != model.StatusSuccess {
					return status, nil
				}
				if data2 != nil {
					data = data2
				}
			} else {
				log.Printf("yaml task name error\n")
				return model.StatusError, nil
			}
		}
	}
	Queue.Done(ctx, gen_id)
	log.Printf("done: %s\n", Queue.Info(ctx))
	return model.StatusSuccess, nil
}
