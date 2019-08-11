package task

import (
	//"fmt"
	//"context"
	"reflect"
	//"errors"

	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

type Task struct {
	Name          string
	Desc          string
	Type          string
	Settings      map[string]interface{}
}

var (
	Funcs = map[string]interface{}{
		// "MQ_Send": MQSend,
		// "HTTP_Send": HTTPSend,
		// "test": Test,
		// "test2": Test2,
		"even":   Even,
		"select": Select,
		// "filter":Filter,
		"accumulate": Accumulate,
	}
)

//map[string]interface{}
func Call(m map[string]interface{}, name string, data []byte, params map[string]interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])

	//if len(params) != f.Type().NumIn() {
	//    err = errors.New("The number of params is not adapted.")
	//    return
	//}
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(data)
	in[1] = reflect.ValueOf(params)
	//for k, param := range params {
	//    in[k] = reflect.ValueOf(param)
	//}
	result = f.Call(in)
	return
}

func CallPipeline(pipeline2 *yaml.Task, data []byte) ([]byte, string) {
	result, _ := Call(Funcs, pipeline2.Type, data, pipeline2.Settings)

	if len(result) >= 2 {
		data = result[0].Interface().([]byte)
		status := result[1].Interface().(string)
		return data, status
	}
	return nil, "success"

	//if pipeline2.Type == "MQ_Send" {
	//	Call(Funcs, pipeline2.Type, pipeline2.Settings)
	//}
	//if pipeline2.Type == "HTTP_Send" {
	//	Call(Funcs, pipeline2.Type, pipeline2.Settings["url"], data)
	//}
	//if pipeline2.Type == "test" {
	//	Call(Funcs, pipeline2.Type, data)
	//}
	//if pipeline2.Type == "test2" {
	//	result, _ := Call(Funcs, pipeline2.Type, data)
	//	data = result[0].Interface().([]byte)
	//	log.Printf("Data :", string(data))
	//	return data
	//}
	//if pipeline2.Type == "split_data" {
	//	result, _ := Call(Funcs, pipeline2.Type, data)
	//	data = result[0].Interface().([]byte)
	//	log.Printf("Data :", string(data))
	//	return data
	//}

	//return nil
}
