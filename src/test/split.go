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
        }else if typeof(data[key]) != "[]interface {}"{
            output_dict[key__+key] = data[key]
        }else{
            for i:=0; i < len(data[key].([]interface{})); i++ {
                dict__,_ := split(data[key].([]interface{})[i].(map[string]interface{}),make(map[string]interface{}),key__ + key + "__")
                output_list = append(output_list, dict__)
            }
        }
    }
    return output_dict,output_list
}

func SplitData(data []byte) ([]byte,error) {
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
    fmt.Println(list_data)

    return data2,nil
}







func main() {
	// setting := "{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-30T07:33:13.161Z\",\"data\":{\"bill_no\":\"string\",\"bill_date\":\"2018-03-03\",\"ops_time\":\"2019-04-30T07:33:13.161Z\",\"details\":[{\"fabric_no\":\"string\",\"line\":\"ASDF\",\"model_no\":\"83234358\",\"model_name\":\"string\",\"item_no\":\"23234543\",\"item_name\":\"string\",\"quantity\":0,\"unit_name\":\"m\",\"order_no\":\"string\",\"order_date\":\"2018-02-10\",\"customer_code\":\"string\",\"customer_name\":\"string\",\"order_delivery_date\":\"2018-03-20\",\"order_quantity\":0,\"width\":0,\"gmwt\":0,\"card_no\":\"string\",\"lot_no\":\"string\",\"sequence_no\":\"string\",\"grade\":\"string\",\"location_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"
	// mySetting:=make(map[string]interface{})
    // json.Unmarshal([]byte(setting),&mySetting)

    str := "{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-30T07:33:13.161Z\",\"data\":{\"bill_no\":\"string\",\"bill_date\":\"2018-03-03\",\"ops_time\":\"2019-04-30T07:33:13.161Z\",\"details\":[{\"fabric_no\":\"string\",\"line\":\"ASDF\",\"model_no\":\"83234358\",\"model_name\":\"string\",\"item_no\":\"23234543\",\"item_name\":\"string\",\"quantity\":0,\"unit_name\":\"m\",\"order_no\":\"string\",\"order_date\":\"2018-02-10\",\"customer_code\":\"string\",\"customer_name\":\"string\",\"order_delivery_date\":\"2018-03-20\",\"order_quantity\":0,\"width\":0,\"gmwt\":0,\"card_no\":\"string\",\"lot_no\":\"string\",\"sequence_no\":\"string\",\"grade\":\"string\",\"location_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"

    bb,err := SplitData([]byte(str))
    fmt.Println(bb,err)
}