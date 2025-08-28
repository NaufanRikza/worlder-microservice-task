package entity

import (
	"time"

	"gorm.io/gorm"
)

type DefaultAttribute struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at"`
}
