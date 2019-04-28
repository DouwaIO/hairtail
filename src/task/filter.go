package task

import (
	"fmt"
    // "github.com/DouwaIO/hairtail/src/schema"
    "encoding/json"
    "reflect"
	"errors"
	"strings"
	"strconv"
)


func Filter(data []byte, params map[string]interface{}) ([]byte,[]byte,error){
	var list_data []map[string]interface{}
	err := json.Unmarshal(data,&list_data)
	if err != nil{
		return nil,nil,err
	}
	
	filter := params["filter"]
	if filter == nil{
		return nil,nil,errors.New("filter不存在")
	}

	// fmt.Println(len(include.([]interface{})))
	// 将[a = b, c = d] 转换为 {"a":"b","c":"d"}

	// filter_map := setting["filter"]

	filter_map := make(map[string]string)

	for i:=0; i<len(filter.([]interface{})); i++{
		d := strings.Split(filter.([]interface{})[i].(string), "=")
		filter_map[d[0]] = d[1]
	}

	fmt.Println(filter_map)

	var res_matching_list []map[string]interface{}
	var res_mismatching_list []map[string]interface{}
	for i:=0; i<len(list_data); i++ {
		flag := 0
		for key := range filter_map {
			// fmt.Println(reflect.TypeOf(list_data[i][key]))
			val_string := ""
			val := list_data[i][key]
			switch val.(type) {
				case string:
					val_string = val.(string)
						// fmt.println("这是一个string类型的变量")
				case int64:
					val_string = strconv.FormatInt(val.(int64),10)
				case float32:
					val_string = fmt.Sprintf("%g", val.(float64))
					// fmt.Print(6666)
				case float64:
					val_string = fmt.Sprintf("%g", val.(float64))
			}

		

			if val_string != filter_map[key]{
				fmt.Println(reflect.TypeOf(list_data[i][key]))
				// fmt.Println(list_data[i][key],"------",filter_map[key])
				res_mismatching_list = append(res_mismatching_list,list_data[i])
				flag = 1
				break
			}
		}
		if flag == 0{
			res_matching_list = append(res_matching_list,list_data[i])
		}
		
	}

	res_matching_byte,_ := json.Marshal(res_matching_list)
	res_mismatching_byte,_ := json.Marshal(res_mismatching_list)

	return res_matching_byte,res_mismatching_byte,nil
	// // return nil
}