package task

import (
	"fmt"
	// "time"
	"strconv"
	"strings"
	//"errors"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	// "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/utils/gid"
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
	// ignore := params.Settings["ignore"].(bool)

	db := params.DB

	for _, r1i := range d1 {
		// log.Debugf("start deal")

		r1 := r1i.(map[string]interface{})
		// log.Debugf("r1: %s", r1)

		key := ""
		for _, m := range maps {
			f0 := strings.Split(m.(string), "=")
			f1 := f0[0]
			f2 := f0[1]

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

			// set source value to target for same column
			r1[f2] = r1[f1]
			key += fmt.Sprintf("%s=%s,", f1, v1s)
		}
		key = strings.TrimRight(key, ",")
		// log.Debugf("key: %s", key)

		// get new row 1
		r1Json, err := json.Marshal(r1)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("marshal r1 error")
			return nil, err
		}
		var d2 model.RemoteData

		sql := `
insert into remote_data(id, key, data, created_at, updated_at)
values (?, ?, ?, current_timestamp, current_timestamp)
on conflict (key)
do update
set updated_at = current_timestamp
returning id, key, data, created_at, updated_at
`
		id := gid.New().String()
		err = db.Raw(sql, id, key, r1Json).Scan(&d2).Error
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("first or create error")
			return nil, err
		}
		// log.Infof("data: %s", d2)

		if d2.CreatedAt != d2.UpdatedAt {
			d2Data, err := d2.Data.Value()
			if err != nil {
				log.Errorf("get data value error: %s", err)
				return nil, err
			}

			var r2 map[string]interface{}
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
