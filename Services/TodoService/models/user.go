package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model

    Username 	string	`gorm:"unique_index"`
    Firstname   string
    Lastname	string

	Todos       []Todo  `gorm:"foreignKey:UserID"`
}
