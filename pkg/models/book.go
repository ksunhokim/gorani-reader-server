package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      string
	BookPicture BookPicture
}

type BookPicture struct {
	gorm.Model
	BookID uint
	URL    string `gorm:"not null"`
}
