package task

import (
	"log"
	"fmt"
	"context"
	"reflect"
	"errors"
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task/queue"
)

var (
	Funcs = map[string]interface{}{"MQ_Send": MQSend,
				       "HTTP_Send": HTTPSend,
				       "test": Test,
					   "test2": Test2,
					   "split_data" :SplitData}
)


func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}

func CallPipeline(pipeline2 *yaml_pipeline.Container, data []byte) []byte {
	s := queue.New()
	ctx := context.Background()
	fmt.Println("s", s)
	log.Printf("start: %s\n", s.Info(ctx))
	task := &queue.Task{
		ID: "1",
		Data: []byte("aaa"),
	}
	s.Push(ctx, task)
	log.Printf("push: %s\n", s.Info(ctx))
	fn, err := queue.CreateFilterFunc("filter")
	if err != nil {
	   fmt.Println("s", err)
	}
	s.Poll(ctx, fn, "abc2")
	log.Printf("poll: %s\n", s.Info(ctx))
	res := s.Extend(ctx, "abc")
	log.Printf("extend err: %s\n", res)
	log.Printf("extend: %s\n", s.Info(ctx))
	s.Done(ctx, "abc")
	log.Printf("done: %s\n", s.Info(ctx))

	if pipeline2.Type == "MQ_Send" {
		Call(Funcs, pipeline2.Type, pipeline2.Settings["protocol"], pipeline2.Settings["host"], pipeline2.Settings["user"], pipeline2.Settings["pwd"], pipeline2.Settings["topic"], data)
	}
	if pipeline2.Type == "HTTP_Send" {
		Call(Funcs, pipeline2.Type, pipeline2.Settings["url"], data)
	}
	if pipeline2.Type == "test" {
		Call(Funcs, pipeline2.Type, data)
	}
	if pipeline2.Type == "test2" {
		result, _ := Call(Funcs, pipeline2.Type, data)
		data = result[0].Interface().([]byte)
		log.Printf("Data :", string(data))
		return data
	}
	if pipeline2.Type == "split_data" {
		result, _ := Call(Funcs, pipeline2.Type, data)
		data = result[0].Interface().([]byte)
		log.Printf("Data :", string(data))
		return data
	}
	return nil
}



