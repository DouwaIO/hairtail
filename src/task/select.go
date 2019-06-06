package task

import (
    "encoding/json"
	"strings"
	"log"
	"time"
)


func SelectData(data []byte, params map[string]interface{}) ([]byte,string){
	log.Println("SelectData")
	
	start := time.Now().Unix()
    log.Println("select start is ",start)
    start += 1


	// log.Println("data is %s",string(data))
	var list_data []map[string]interface{}
	err := json.Unmarshal(data,&list_data)
	if err != nil{
		log.Printf("%s", err)
		return nil,"error"
	}
	
    include := params["include"]
	exclude := params["exclude"]
	rename := params["rename"]

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
		}else{
			res_list = list_data
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
	
	res_byte,_ := json.Marshal(res_list)
	return res_byte,"success"
	// return nil
}







