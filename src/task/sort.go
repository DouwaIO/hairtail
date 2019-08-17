package task

import (
	"encoding/json"
	"sort"

	log "github.com/sirupsen/logrus"
)

func Sort(params *Params) (*Result, error) {
    // log.Debugf("sort data: %s", params.Data)

	var data []map[string]interface{}
	err := json.Unmarshal(params.Data, &data)
	if err != nil {
		log.Errorf("task sort unmarshal data error: %s", err)
		return nil, err
	}

	keys := params.Settings["keys"].([]interface{})

	sort.Slice(data, func(i, j int) bool {
        less := false
        for _, key := range keys {
            k := key.(string)

            v1 := data[i][k].(string)
            v2 := data[j][k].(string)

            if v1 > v2 {
                break
            } else if v1 < v2 {
                less = true
                break
            } else {
                continue
            }
        }
		return less
	})

	data2, _ := json.Marshal(data)
    // log.Debugf("sorted data: %s", data2)
    result := Result{
        Data: data2,
    }
	return &result, nil
}
