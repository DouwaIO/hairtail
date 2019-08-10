package model

type Task struct {
	BaseModel
	Name       string `json:"name"          gorm:"type:varchar(128);column:name"`
	Type       string `json:"type"          gorm:"type:varchar(128);column:type"`
	Desc       string `json:"desc"          gorm:"type:varchar(500);column:desc"`
	WorkFlowId string `json:"workflow_id"          gorm:"type:varchar(50);column:workflow_id"`
	Settings   string `json:"settings"`
}
