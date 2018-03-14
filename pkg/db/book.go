package db

import "github.com/jinzhu/gorm"

type (
	Book struct {
		gorm.Model
		Title       string `gorm:"not null"`
		Author      string
		BookPicture BookPicture
	}
	BookPicture struct {
		gorm.Model
		BookID uint   `gorm:"not null"`
		URL    string `gorm:"not null"`
	}
)
