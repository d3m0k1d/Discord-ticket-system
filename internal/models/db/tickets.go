package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title" validate:"required,min=3,max=255"`
	Description string         `gorm:"type:text;not null" json:"description" validate:"required"`
	Status      bool           `gorm:"not null" json:"status"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
