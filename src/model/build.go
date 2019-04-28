package model

type Build struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:build_id"`
	Service string `json:"service"          gorm:"type:varchar(50);column:service_id"`
	Data  string `json:"data"             gorm:"type:text;column:data""`
	Status string `json:"status"    gorm:"column:status"`
	Message   string  `json:"message"   gorm:"column:message"`
	Timestamp int64   `json:"timestamp"  gorm:"column:timestamp"`
	//DataID string `json:"data_id"             gorm:"column:data_id"`
}
