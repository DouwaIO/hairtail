package model

type RemoteData struct {
	ID string `json:"id"             gorm:"primary_key;type:varchar(50);column:id"`
	Name string `json:"name"          gorm:"type:varchar(500);column:name"`
	// Data string `json:"data"          gorm:"type:text;column:data"`
	Data []byte `json:"data"      gorm:"column:data"`
}


func (RemoteData) TableName() string {
	return "remote_data"
  }