package main

import (
    "fmt"
    "encoding/json"
//    "strconv"
)


func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}


func split(data map[string]interface{},output_dict map[string]interface{},key__ string) (map[string]interface{}, []interface{}){
    var output_list []interface{}
    var output_list2 []interface{}
    var output_list3 []interface{}
    for key := range data{
        if typeof(data[key]) == "map[string]interface {}" {
	    fmt.Println("map", data[key])
            _,list__ := split(data[key].(map[string]interface{}),output_dict,key__ + key + "__")
            output_list = append(output_list, list__...)
	    fmt.Println("out map", output_list)
        }else if typeof(data[key]) != "[]interface {}"{
	    fmt.Println("str", data[key].(string))
            output_dict[key__+key] = data[key]
        }else{
	    fmt.Println("list", data[key])
            for i:=0; i < len(data[key].([]interface{})); i++ {
                dict__,_ := split(data[key].([]interface{})[i].(map[string]interface{}),make(map[string]interface{}),key__ + key + "__")
                //output_list = append(output_list, dict__)
		if len(output_list2) == 0 {
                  output_list = append(output_list, dict__)
	        }

            }

	    if len(output_list2) > 0 {
            output_list2 = data[key].([]interface{})
	    for _, v := range output_list {
	        for k, v3 := range v.(map[string]interface{}) {
	           for _, v2 := range output_list2 {
	               for k2, v4 := range v2.(map[string]interface{}) {
			    qq := make(map[string]interface{})
			    qq[k2] = v4
			    qq[k] = v3
	                    output_list3 = append(output_list3, qq)
			    }
			}
		   }
	        }
             }
            output_list2 = data[key].([]interface{})
	    if len(output_list3) != 0 {
		    output_list = output_list3
	    }
       }

    }

    fmt.Println(len(output_list))
    return output_dict,output_list
}

func split_data(data []byte) []byte {
    myMap:=make(map[string]interface{})
    json.Unmarshal(data,&myMap)
    output_dict := make(map[string]interface{})
    dict_data,list_data := split(myMap,output_dict,"")

    if len(list_data) == 0{
        list_data = append(list_data,dict_data)
    }else{
        for i:=0; i < len(list_data); i++ {
            for k ,v := range dict_data { list_data[i].(map[string]interface{})[k] = v }
        }
    }

    data2, _ := json.Marshal(list_data)

    return data2
}




func main() {

    str := "{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]},\"customers2\":[{\"name2\":\"安踏2\"},{\"name2\":\"李宁2\"}]}"
    //str := "{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}],\"customers2\":[{\"name2\":\"安踏2\"},{\"name2\":\"李宁2\"}]}}"



    bb := split_data([]byte(str))
    fmt.Println(string(bb))
}
