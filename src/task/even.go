package task

import (
	"fmt"
	"log"
	// "github.com/DouwaIO/hairtail/src/schema"
	"encoding/json"
	// "reflect"
	// "errors"
	"time"
	// "log"
)

var start int

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func split(data map[string]interface{}, output_dict map[string]interface{}, key__ string) (map[string]interface{}, []interface{}) {
	var output_list []interface{}
	for key := range data {
		if typeof(data[key]) == "map[string]interface {}" {
			_, list__ := split(data[key].(map[string]interface{}), output_dict, key__+key+"__")
			output_list = append(output_list, list__...)
		} else if typeof(data[key]) != "[]interface {}" {
			output_dict[key__+key] = data[key]
		} else {
			for i := 0; i < len(data[key].([]interface{})); i++ {
				dict__, _ := split(data[key].([]interface{})[i].(map[string]interface{}), make(map[string]interface{}), key__+key+"__")
				output_list = append(output_list, dict__)
			}
		}
	}
	return output_dict, output_list
}

func Even(data []byte, params map[string]interface{}) ([]byte, string) {
	start := time.Now().Unix()
	log.Println("split_data start is ", start)
	start += 1

	log.Println("SplitData")
	myMap := make(map[string]interface{})
	json.Unmarshal(data, &myMap)
	output_dict := make(map[string]interface{})
	dict_data, list_data := split(myMap, output_dict, "")

	if len(list_data) == 0 {
		list_data = append(list_data, dict_data)
	} else {
		for i := 0; i < len(list_data); i++ {
			for k, v := range dict_data {
				list_data[i].(map[string]interface{})[k] = v
			}
		}
	}

	data2, _ := json.Marshal(list_data)
	// fmt.Println(string(data2))

	// log.Println("result", string(data2))
	return data2, "success"
}
