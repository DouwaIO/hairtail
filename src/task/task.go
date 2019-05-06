package task

import (
	//"fmt"
	//"context"
	"reflect"
	//"errors"
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task/queue"
	//"github.com/DouwaIO/hairtail/src/utils"
	// "log"
)

var Queue queue.Queue

var (
	Funcs = map[string]interface{}{"MQ_Send": MQSend,
				       "HTTP_Send": HTTPSend,
				       "test": Test,
					   "test2": Test2,
					   "even": SplitData,
					   "select": SelectData,
					   "filter":Filter,
					   "accumulate":Accumulate}
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

func CallPipeline(pipeline2 *yaml_pipeline.Container, data []byte) []byte {
	result,_ := Call(Funcs, pipeline2.Type, data, pipeline2.Settings)

	if len(result) > 0 {
		data = result[0].Interface().([]byte)
		return data
	}
	return nil



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

//func CallPipeline(pipeline2 *yaml_pipeline.Container, data []byte) []byte {
//	//s := queue.New()
//	//ctx := context.Background()
//	//fmt.Println("s", Queue)
//	////log.Printf("start: %s\n", Queue.Info(ctx))
//	//gen_id := utils.GeneratorId()
//	//task := &queue.Task{
//	//	ID: gen_id,
//	//	Data: data,
//	//}
//	//Queue.Push(ctx, task)
//
//	log.Printf("push: %s\n", Queue.Info(ctx))
//
//	//res := Queue.Extend(ctx, "abc")
//	//log.Printf("extend err: %s\n", res)
//	//log.Printf("extend: %s\n", Queue.Info(ctx))
//	//Queue.Done(ctx, "abc")
//	//log.Printf("done: %s\n", Queue.Info(ctx))
//
//	if pipeline2.Type == "MQ_Send" {
//		Call(Funcs, pipeline2.Type, pipeline2.Settings["protocol"], pipeline2.Settings["host"], pipeline2.Settings["user"], pipeline2.Settings["pwd"], pipeline2.Settings["topic"], data)
//	}
//	if pipeline2.Type == "HTTP_Send" {
//		Call(Funcs, pipeline2.Type, pipeline2.Settings["url"], data)
//	}
//	if pipeline2.Type == "test" {
//		Call(Funcs, pipeline2.Type, data)
//	}
//	if pipeline2.Type == "test2" {
//		result, _ := Call(Funcs, pipeline2.Type, data)
//		data = result[0].Interface().([]byte)
//		log.Printf("Data :", string(data))
//		return data
//	}
//	if pipeline2.Type == "split_data" {
//		result, _ := Call(Funcs, pipeline2.Type, data)
//		data = result[0].Interface().([]byte)
//		log.Printf("Data :", string(data))
//		return data
//	}
//
//	fn, err := queue.CreateFilterFunc("filter")
//	if err != nil {
//	   fmt.Println("s", err)
//	}
//	Queue.Poll(ctx, fn, gen_id)
//	log.Printf("poll: %s\n", Queue.Info(ctx))
//
//	return nil
//}
//
//
//
