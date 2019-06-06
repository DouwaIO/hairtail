package pipeline

import (
	"context"
	"log"
	"fmt"
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
	"github.com/DouwaIO/hairtail/src/pipeline/queue"
	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/utils"
)

var Queue queue.Queue


func Pipeline(pipeline []*yaml_pipeline.Container, data []byte) string {
	//parsed, err := yaml_pipeline.ParseString(q.config)
	//if err != nil {
	//	return errors.New("yaml type error")
	//}
	log.Printf("hello world")
        ctx := context.Background()
        gen_id := utils.GeneratorId()
        task2 := &queue.Task{
	         ID: gen_id,
	         Data: data,
        }
        Queue.Push(ctx, task2)
        fn, err := queue.CreateFilterFunc("filter")
        if err != nil {
           fmt.Println("s", err)
        }
	Queue.Poll(ctx, fn, gen_id)
	// log.Printf("poll: %s\n", Queue.Info(ctx))

	if len(pipeline) > 0 {
		for _, pipeline2 := range pipeline {
			if _, ok := task.Funcs[pipeline2.Type]; ok {
				data2, status := task.CallPipeline(pipeline2, data)
				if status != model.StatusSuccess {
					return status
				}
				if data2 != nil {
					data = data2
				}
			} else {
				log.Printf("yaml task name error\n")
				return model.StatusError
			}

		}
	}
	Queue.Done(ctx, gen_id)
	log.Printf("done: %s\n", Queue.Info(ctx))
	return model.StatusSuccess
}


