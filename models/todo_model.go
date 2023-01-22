package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ActivityGroupId uint           `json:"activity_group_id"`
	Title           string         `json:"title"`
	IsActive        bool           `gorm:"default:true" json:"is_active"`
	Priority        string         `gorm:"default:very-high" json:"priority"`
}

type CreateTodoRequest struct {
	Title           string `validate:"required" json:"title"`
	ActivityGroupId uint   `validate:"required" json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type UpdateTodoRequest struct {
	Title           string `validate:"required" json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}
