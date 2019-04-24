package task

import (
	"log"
	"reflect"
	"errors"
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

var (
	Funcs = map[string]interface{}{"MQ_Send": MQSend,
				       "HTTP_Send": HTTPSend,
				       "test": Test,
				       "test2": Test2 }
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
	return nil
}

