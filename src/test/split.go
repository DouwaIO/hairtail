package main

import (
	"fmt"
    // "github.com/DouwaIO/hairtail/src/schema"
    "encoding/json"
    // "reflect"
    // "errors"
)


func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}


func split(data map[string]interface{},output_dict map[string]interface{},key__ string) (map[string]interface{}, []interface{}){
    var output_list []interface{}
    for key := range data{
        if typeof(data[key]) == "map[string]interface {}"{
            _,list__ := split(data[key].(map[string]interface{}),output_dict,key__ + key + "__")
            output_list = append(output_list, list__...)
            fmt.Println("out_",list__)
        }else if typeof(data[key]) != "[]interface {}"{
            output_dict[key__+key] = data[key]
        }else{
            for i:=0; i < len(data[key].([]interface{})); i++ {
                dict__,_ := split(data[key].([]interface{})[i].(map[string]interface{}),make(map[string]interface{}),key__ + key + "__")
                fmt.Println(dict__)
                output_list = append(output_list, dict__)
                fmt.Println(output_list)
            }
        }
        
    }

    fmt.Println("out",output_list)
    return output_dict,output_list
}




func main() {

    str := "{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-23T03:10:33.775Z\",\"data\":{\"bill_no\":\"string\",\"details\":[{\"grade\":\"string\",\"location_no\":\"string\",\"line\":\"ASDF111\"},{\"fabric_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"
    myMap:=make(map[string]interface{})
    output_dict := make(map[string]interface{})
      
    json.Unmarshal([]byte(str),&myMap)

    aa,bb := split(myMap,output_dict,"")
    fmt.Println(aa)
    fmt.Println("test",bb)
    fmt.Println(output_dict)

}
