package model

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type RemoteData struct {
	BaseModel
	Key string `json:"key"          gorm:"type:varchar(500);column:key;unique"`
	Data postgres.Jsonb `json:"data"      gorm:"column:data"`
}

func (RemoteData) TableName() string {
	return "remote_data"
}
