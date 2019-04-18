package model

type Schema struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:schema_id"`
	Name string `json:"_"          gorm:"type:varchar(500);column:name"`

	Data string `json:"data"          gorm:"type:text;column:data"`
}

