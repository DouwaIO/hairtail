package main

import (
    "fmt"
    "encoding/json"
//    "strconv"
)


func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func mergeList(listA []interface{}, key string, listB []interface{}) []interface{} {
    fmt.Printf("merge list: %s = %+v and %+v\n", key, listA, listB)
    var newList []interface{}
    if len(listB) > 0 {
        for _, cR := range listA {
            for _, nR := range listB {
                dict := make(map[string]interface{})
                for nRK, nRV := range nR.(map[string]interface{}) {
                    dict[nRK] = nRV
                }
                for cRK, cRV := range cR.(map[string]interface{}) {
                    if key != "" {
                      dict[fmt.Sprintf("%s__%s", key, cRK)] = cRV
                    } else {
                      dict[cRK] = cRV
                    }
                }
                newList = append(newList, dict)
            }
        }
    } else {
        for _, cR := range listA {
            dict := make(map[string]interface{})
            // fmt.Printf("%+v\n", typeof(cR))
            for cRK, cRV := range cR.(map[string]interface{}) {
                if key != "" {
                  dict[fmt.Sprintf("%s__%s", key, cRK)] = cRV
                } else {
                  dict[cRK] = cRV
                }
            }
            newList = append(newList, dict)
        }
    }
    return newList
}

func mergeObj(listA []interface{}, obj map[string]interface{}) []interface{} {
    fmt.Printf("merge object: %+v and %+v\n", listA, obj)
    if len(listA) > 0 {
        for _, r := range listA {
            for k, v := range obj {
                r.(map[string]interface{})[k] = v
            }
        }
    } else {
        listA = append(listA, obj)
    }
    return listA
}

func even(rows []interface{}, level int) ([]interface{}){
    level++
    var rstList []interface{}
    // 遍历所有行
    for _, r := range rows {
        rstObj := make(map[string]interface{})
        // 遍历所有key
        for k, v := range r.(map[string]interface{}) {
            // fmt.Printf("level: %d, key: %s = %+v\n", level, k, v)
            // 如果是对象
            if typeof(v) == "map[string]interface {}" {
                var list []interface{}
                list = append(list, v.(map[string]interface{}))
                rstObj[k] = even(list, level)
                fmt.Printf("level: %d, dict: %+v\n", level, rstObj[k])
            // 如果是列表
            } else if typeof(v) == "[]interface {}" {
                rstObj[k] = even(v.([]interface{}), level)
                fmt.Printf("level: %d, list: %+v\n", level, rstObj[k])
            // 如果是常量
            } else {
                rstObj[k] = v
                // fmt.Printf("level: %d, const: %s = %s\n", level, k, v)
            }
        }

        // fmt.Printf("level: %d, only obj: %+v\n", level, rstObj)

        var rowList []interface{}
        for k, v := range rstObj {
            if typeof(v) == "[]interface {}" {
                rowList = mergeList(v.([]interface{}), k, rowList)
                delete(rstObj, k)
            }
        }
        // fmt.Printf("level: %d, only list: %+v\n", level, rowList)
        rowList = mergeObj(rowList, rstObj)

        fmt.Printf("level: %d, all list: %+v\n", level, rowList)
        if len(rstList) == 0 {
            rstList = rowList
        } else {
            rstList = append(rstList, rowList...)
        }
    }
    return rstList
}

func split_data(data []byte) []byte {
    var myList []interface{}
    json.Unmarshal(data, &myList)

    level := 0
    evenList := even(myList, level)

    // for i:=0; i < len(list_data); i++ {
    //     for k ,v := range dict_data { list_data[i].(map[string]interface{})[k] = v }
    // }

    data2, _ := json.Marshal(evenList)
    return data2
}




func main() {

    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}]}}]"
    str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}]},\"data1\":{\"order_no\":\"001\",\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"},{\"name\":\"特步\"}]}}]"
    // str := "[{\"datetime\":\"2019-04-03 10:22:20\",\"data\":{\"order_no\":\"001\",\"sales\":[{\"name\":\"张三\"},{\"name\":\"李四\"}],\"customers\":[{\"name\":\"安踏\"},{\"name\":\"李宁\"},{\"name\":\"特步\"}]}}]"

    fmt.Printf("data: %s\n\n", str)
    bb := split_data([]byte(str))
    fmt.Printf("\nresult: \n%s\n", string(bb))
}
