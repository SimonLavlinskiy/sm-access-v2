package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string     `sql:"type:uuid;primary_key; default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"softDelete" json:"deleted_at"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	scope.Statement.SetColumn("ID", id)
	return nil
}
