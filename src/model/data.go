package model

type Data struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:data_id"`
	Service string `json:"service_id"             gorm:"primary_key;type:varchar(50);column:service_id"`
	Data string `json:"data"          gorm:"type:text;column:data"`
}
