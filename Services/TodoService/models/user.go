package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model

	ID	        uint	`gorm:"primary_key"`
    Username 	string	`gorm:"unique_index"`
    Firstname   string
    Lastname	string

	Todos       []Todo  `gorm:"foreignKey:UserID"`
}
