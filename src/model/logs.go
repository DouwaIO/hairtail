package model


type LogData struct {
	ID     string  `meddler:"log_id,pk"    gorm:"primary_key;type:varchar(50);column:log_id"`
	ProcID string  `meddler:"log_job_id"   gorm:"type:varchar(50);column:log_job_id"`
	Data   string `meddler:"log_data"     gorm:"type:text;column:data"`
}

// 定义生成表的名称
func (LogData) TableName() string {
	return "logs"
  }
