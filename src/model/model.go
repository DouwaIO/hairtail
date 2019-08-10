package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/DouwaIO/hairtail/src/utils/gid"
)

type BaseModel struct {
    ID        string    `json:"id" gorm:"type:varchar(50);primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at"`
	Extend postgres.Jsonb `json:"extend"`
}

func (baseModel *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", gid.New().String())
	return nil
}
