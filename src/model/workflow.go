package model

type Workflow struct {
	BaseModel
	Name     string `json:"name"          gorm:"type:varchar(128);column:name"`
	Config   string `json:"config"`
	Activate int    `json:"activate"`
}
