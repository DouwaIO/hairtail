package model

type Data struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:data_id"`
	Type string `json:"type"          gorm:"type:varchar(500);column:type"`
	Name string `json:"name"          gorm:"type:varchar(500);column:name"`
	Data string `json:"data"          gorm:"type:text;column:data"`
}
