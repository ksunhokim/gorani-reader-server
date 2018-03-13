package models

import "github.com/jinzhu/gorm"

type Word struct {
	gorm.Model
	Word   string
	Pron   string
	Source string
	Type   string
	Def    []Def
}

type Def struct {
	gorm.Model
	WordID  uint
	Part    string
	Def     string
	Example []Example
}

type Example struct {
	gorm.Model
	DefID uint
	Kor   string
	Eng   string
}
