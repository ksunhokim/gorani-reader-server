package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Books    []Book `gorm:"many2many:user_books;"`
}
