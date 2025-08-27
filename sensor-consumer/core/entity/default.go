package entity

import "time"

type DefaultAttribute struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}
