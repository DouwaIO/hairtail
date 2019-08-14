package task

import (
	// "fmt"
	// "context"
	"reflect"
	// "errors"
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

func (t *Task) Run(data []byte) ([]byte, error) {
	fn := reflect.ValueOf(Funcs[t.Type])

	myMap := make(map[string]interface{})
	json.Unmarshal(data, &myMap)

	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(data)
	in[1] = reflect.ValueOf(t.Settings)
    result := fn.Call(in)

	if len(result) >= 2 {
        resultData := result[0].Interface().([]byte)
		err        := result[1].Interface().(error)
		return resultData, err
	}
	return nil, nil
}
