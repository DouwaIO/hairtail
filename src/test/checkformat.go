package main

import (
	"fmt"
    "github.com/DouwaIO/hairtail/src/schema"
    "encoding/json"
    "reflect"
    "errors"
)

func check(schema *schema.Schema, data []byte) (map[string]interface{},error){
    m := make(map[string]interface{})
    retult := make(map[string]interface{})
    err := json.Unmarshal(data, &m)
    if err != nil{
        return nil,err
    }
    
    column := schema.Columns
    for i:=0; i<len(column); i++ {
        // 判断值是否为空 
        if m[column[i].Name] == nil{
            retult[column[i].Name] = column[i].Default
            continue
        }
        // 判断类型是否正确
        column_type := reflect.TypeOf(m[column[i].Name]).Name()
        if column_type == "float64"{
            column_type = "int"
        }
        if column_type != column[i].Type{
            return nil,errors.New(column[i].Name+"类型错误")
        }else{
            retult[column[i].Name] = m[column[i].Name]
        }

    }
    return retult,nil
}

func main() {
test := `version: v1
kind: model
name: storage
columns:
  - name: id
    desc: a identity data
    type: int
    default: 0
  - name: name
    desc: the name of people
    type: string
    default: ''
  - name: sex
    desc: '0: male, 1: female'
    type: int
    default: 0
    `
	branches, err := schema.ParseString(test)
	if err != nil {
		fmt.Println(err)
	}
    // fmt.Println(branches)
    

    json_str := `{"id":1,"name": "jjq","sex": 1, "aa":1}`
    // m := make(map[string]interface{})
    // err = json.Unmarshal([]byte(json_str), &m)
    result,err := check(branches,[]byte(json_str))
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println(result)
    }

    // fmt.Println(m["name"])

}
