package db

import "github.com/jinzhu/gorm"

type (
	Word struct {
		gorm.Model
		Word   string `gorm:"unique;not null"`
		Pron   string
		Source string `gorm:"unique;not null"`
		Type   string `gorm:"not null"`
		Def    []Def
	}

	Def struct {
		gorm.Model
		WordID  uint `gorm:"not null"`
		Part    string
		Def     string `gorm:"not null"`
		Example []Example
	}

	Example struct {
		gorm.Model
		DefID uint `gorm:"not null"`
		Kor   string
		Eng   string
	}
)
