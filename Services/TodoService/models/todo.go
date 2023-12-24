package models

import (
	"gorm.io/gorm"
)

type Todo struct {
    gorm.Model

    Priority    uint    `gorm:"index;default:3"`
    Title       string
    Description string
    Status      string  `gorm:"index;default:open"`

    UserID uint         `gorm:"not null"`
}
