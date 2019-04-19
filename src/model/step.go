package model


type Step struct {
	ID string `json:"id"    gorm:"primary_key;type:varchar(50);column:log_id"`
	Build string  `meddler:"build_id"   gorm:"type:integer;column:build_id"`
	Num int64   `json:"timestamp"  gorm:"column:num"`
	Status string `json:"status"    gorm:"column:status"`
	//Data   []byte `meddler:"log_data"     gorm:"type:mediumblob;column:log_data"`
}
