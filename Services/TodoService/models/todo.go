package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
    gorm.Model

    ID          uint       `gorm:"primary_key"`
    Priority    uint       `gorm:"index"`
    Title       string
    Description string
    Status      string      `gorm:"index"`

    CreatedAt   time.Time   `gorm:"not null"`
    UpdatedAt   time.Time   `gorm:"not null"`

    UserID      uint        `gorm:"not null"`
}
