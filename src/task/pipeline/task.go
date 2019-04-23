package pipeline

import (
	"log"
	"reflect"
	"errors"
	"github.com/DouwaIO/hairtail/src/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
)

var (
	funcs = map[string]interface{}{"MQ_Send": MQSend,
				       "HTTP_Send": HTTPSend,
				       "test":Test,
				       "test2":Test2 }
)

type fifo struct {
	//sync.Mutex

	//workers   map[*worker]struct{}
	//running   map[string]*entry
	//pending   *list.List
	//extension time.Duration
	//config string
	data []byte
	pipeline []*pipeline.Container
}

func New(pipeline2 []*pipeline.Container, data []byte) task.Pipeline  {
	return &fifo{
		pipeline: pipeline2,
		data: data,
	}
}


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

func Call_func(pipeline2 *pipeline.Container, data []byte) []byte {
	if pipeline2.Type == "MQ_Send" {
		Call(funcs, pipeline2.Type, pipeline2.Settings["protocol"], pipeline2.Settings["host"], pipeline2.Settings["user"], pipeline2.Settings["pwd"], pipeline2.Settings["topic"], data)
	}
	if pipeline2.Type == "HTTP_Send" {
		Call(funcs, pipeline2.Type, pipeline2.Settings["url"], data)
	}
	if pipeline2.Type == "test" {
		Call(funcs, pipeline2.Type, data)
	}
	if pipeline2.Type == "test2" {
		result, _ := Call(funcs, pipeline2.Type, data)
		data = result[0].Interface().([]byte)
		log.Printf("Data :", string(data))
		return data
	}
	return nil
}

func (q *fifo) Pipeline() error {
	//parsed, err := pipeline.ParseString(q.config)
	//if err != nil {
	//	return errors.New("yaml type error")
	//}
	if len(q.pipeline) > 0 {
		for _, pipeline2 := range q.pipeline {
			if _, ok := funcs[pipeline2.Type]; ok {
				data := Call_func(pipeline2, q.data)
				if data != nil {
					q.data = data
				}
			} else {
				return nil
			}

		}
	}
	return nil
}


