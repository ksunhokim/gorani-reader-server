package models

type Word struct {
	Id            int    `gorm:"column:word_id;primary_key"`
	Word          string `gorm:"column:word"`
	Pronunciation string `gorm:"column:word_pronunciation"`
}
