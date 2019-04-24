package model

type Service struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:service_id"`
	Name string `json:"name"          gorm:"type:varchar(500);column:name"`
	Type string `json:"type"          gorm:"type:varchar(500);column:type"`
	Pipeline string `json:"pipeline_id"          gorm:"type:varchar(50);column:pipeline_id"`
	Data string `json:"data"          gorm:"type:text;column:data"`
	Timestamp int64 `json:"timestamp"          gorm:"column:timestamp"`
}

