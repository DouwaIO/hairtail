package main

import (
	"fmt"
    // "github.com/DouwaIO/hairtail/src/schema"
    "encoding/json"
    // "reflect"
	// "errors"
	"strings"
)




func select_data(data []byte, setting map[string]interface{}) ([]map[string]interface{},error){
	var list_data []map[string]interface{}
	err := json.Unmarshal(data,&list_data)
	if err != nil{
		return nil,err
	}
	
    include := setting["include"]
	exclude := setting["exclude"]
	rename := setting["rename"]

	// fmt.Println(len(include.([]interface{})))
	// 将[a = b, c = d] 转换为 {"a":"b","c":"d"}

	rename_map := make(map[string]string)

	for i:=0; i<len(rename.([]interface{})); i++{
		d := strings.Split(rename.([]interface{})[i].(string), "=")
		rename_map[d[0]] = d[1]
	}

	var res_list []map[string]interface{}
	for i:=0; i<len(list_data); i++ {
		if include != nil{
			// fmt.Println(include.([]interface{})[0])
			map_data:=make(map[string]interface{})
			for j:=0; j<len(include.([]interface{})); j++{
				if list_data[i][include.([]interface{})[j].(string)] != nil{
					map_data[include.([]interface{})[j].(string)] = list_data[i][include.([]interface{})[j].(string)]
				}	
			}
			res_list = append(res_list,map_data)
		}else if exclude != nil{
			for j:=0; j<len(exclude.([]interface{})); j++{
				delete(list_data[i],exclude.([]interface{})[j].(string))
			}
			res_list = append(res_list,list_data[i])
		}
	}



	// 进行重命名
	for i := 0; i<len(res_list); i++{
		for key := range res_list[i]{
			if rename_map[key] != ""{
				res_list[i][rename_map[key]] = res_list[i][key]
				delete(res_list[i],key)
			}

		}
	}
	
	return res_list,nil
	// return nil
}



func main() {
	setting := "{\"include\":[\"a\",\"b\"],\"rename\":[\"b=bbbbb\",\"data__sales__name=sale_name\"]}"
	mySetting:=make(map[string]interface{})
    json.Unmarshal([]byte(setting),&mySetting)

    str := "[{\"a\":1,\"b\":2,\"c\":3}]"

    bb,err := select_data([]byte(str),mySetting)
    fmt.Println(bb,err)
}





