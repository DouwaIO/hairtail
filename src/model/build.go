package model

type Build struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:build_id"`
	Pipeline string `json:"pipeline"          gorm:"type:varchar(50);column:pipeline_id"`
	//DataID string `json:"data_id"             gorm:"column:data_id"`
	Status string `json:"status"    gorm:"column:status"`
	Message   string  `json:"message"   gorm:"column:message"`
	Timestamp int64   `json:"timestamp"  gorm:"column:timestamp"`
}
