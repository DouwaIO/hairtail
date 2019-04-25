package main

import (
    "fmt"
    "encoding/json"
)


func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}


//func split_map(data map[string]interface{}) []interface{} {
//}
func split_list(data []interface{}) []interface{} {

  fmt.Println("list", data)
  return data
}


func split(data map[string]interface{}, key__ string) []interface{} {
    var out_l []interface{}
    out_d := make(map[string]interface{})
    var out_l2 []interface{}
    for key, value := range data {
        if typeof(value) == "string" || typeof(value) == "float64"{
	       out_d[key] = value
               fmt.Println("str", value)
        } else if typeof(value) == "map[string]interface {}"{
               fmt.Println("map", value)
	       fmt.Println("map", data[key])
               list_ := split(value.(map[string]interface{}), key__ + key + "__")
               //output_list = append(output_list, list__...)
	       fmt.Println("out map", list_)
               //list__ := split(value.(map[string]interface{}), key__ + key + "__")
               out_l = append(out_l, list_...)
	       //for k,v := range out_d {

	       //}
        }else if typeof(data[key]) == "[]interface {}"{
		list_ := split_list(data[key].([]interface{}))
		if len(out_l2) != 0 {
		    for _, v := range out_l2 {
			for _, v2 := range out_l {
				for k, v3 := range v.(map[string]interface{}) {
                                    v2.(map[string]interface{})[k] = v3
				}
			}

		    }
		} else {
               	out_l = append(out_l, list_...)
		}
		out_l2 = list_
	        fmt.Println("out map", list_)
        }
    }
    return out_l
}





func main() {

    //str := "{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-23T03:10:33.775Z\",\"data\":{\"bill_no\":\"string\",\"details\":[{\"grade\":\"string\",\"location_no\":\"string\",\"line\":\"ASDF111\"},{\"fabric_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"
    //str := "{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}}"
   str := "{\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"},{\"name\":5}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}, \"datetime\":\"2019-04-03 10:22:20\"}"

    myMap:=make(map[string]interface{})
    json.Unmarshal([]byte(str),&myMap)
    list_data := split(myMap, "")
    fmt.Println("list", list_data)
}
