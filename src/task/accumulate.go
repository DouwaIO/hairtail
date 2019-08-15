package task

import (
	"fmt"
	"strconv"
	"strings"
	//"errors"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	// "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DouwaIO/hairtail/src/model"
)

func Accumulate(params *Params) (*Result, error) {
	var d1 []interface{}
	err := json.Unmarshal(params.Data, &d1)
	if err != nil {
		log.Errorf("data unmarshal error: %s", err)
		return nil, err
	}

	maps := params.Settings["map"].([]interface{})
	source := params.Settings["source"].(string)
	target := params.Settings["target"].(string)
	compute := params.Settings["compute"].(string)
	ignore := params.Settings["ignore"].(bool)

	db := params.DB

	var r2 map[string]interface{}
	for i := 0; i < len(d1); i++ {
		// log.Debugf("start deal")

		r1 := d1[i].(map[string]interface{})
		// log.Debugf("r1: %s", r1)

		key := ""
		for _, m := range maps {
			f0 := strings.Split(m.(string), "=")
			f1 := f0[0]
			// f2 := f0[1]

			v1 := r1[f1]
			v1s := ""
			switch v1.(type) {
			case string:
				v1s = v1.(string)
			case int64:
				v1s = strconv.FormatInt(v1.(int64), 10)
			case float32:
				v1s = fmt.Sprintf("%g", v1.(float64))
			case float64:
				v1s = fmt.Sprintf("%g", v1.(float64))
			}

			key += fmt.Sprintf("%s=%s,", f1, v1s)
		}
		key = strings.TrimRight(key, ",")

		var d2 = new(model.RemoteData)
		err := db.Where("key = ?", key).First(&d2).Error
		// insert
		if err != nil && !ignore {
			for _, m := range maps {
				f0 := strings.Split(m.(string), "=")
				f1 := f0[0]
				f2 := f0[1]

				r1[f2] = r1[f1]
			}
			r1Json, err := json.Marshal(r1)
			if err != nil {
				log.Errorf("marshal r1 error: %s", err)
				return nil, err
			}
			data1 := model.RemoteData{
				Key:  key,
				Data: postgres.Jsonb{r1Json},
			}
			// log.Debugf("create data")
			err = db.Create(&data1).Error
			if err != nil {
				log.Errorf("create data error: %s", err)
				return nil, err
			}
		} else {
			d2Data, err := d2.Data.Value()
			if err != nil {
				log.Errorf("get data value error: %s", err)
				return nil, err
			}

			err = json.Unmarshal(d2Data.([]byte), &r2)
			if err != nil {
				log.Errorf("get unmarshal data error: %s", err)
				return nil, err
			}
			// log.Debugf("r2: %s", r2)

			switch compute {
			case "+":
				r2[target] = r2[target].(float64) + r1[source].(float64)
			case "-":
				r2[target] = r2[target].(float64) - r1[source].(float64)
			case "*":
				r2[target] = r2[target].(float64) * r1[source].(float64)
			case "/":
				r2[target] = r2[target].(float64) / r1[source].(float64)
			}

			r2Json, err := json.Marshal(r2)
			if err != nil {
				log.Errorf("marshal r2 error: %s", err)
				return nil, err
			}
			d2.Data = postgres.Jsonb{r2Json}

			// log.Debugf("save data")
			err = db.Save(&d2).Error
			if err != nil {
				log.Errorf("save data error: %s", err)
				return nil, err
			}
		}
	}

	if err != nil {
		log.Println(err)
	}

	return nil, nil
}
