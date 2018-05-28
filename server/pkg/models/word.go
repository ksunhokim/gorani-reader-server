package models

import "github.com/jinzhu/gorm"

type Word struct {
	Id            int     `gorm:"column:word_id;primary_key"`
	Word          string  `gorm:"column:word"`
	Pronunciation *string `gorm:"column:word_pronunciation"`
	Definitions   []Definition
}

func (Word) TableName() string {
	return "word"
}

type Definition struct {
	Id         int    `gorm:"column:definition_id;primary_key"`
	WordId     int    `gorm:"column:word_id"`
	Definition string `gorm:"column:definition"`
	POS        *int   `gorm:"column:definition_pos"`
	Examples   []Example
}

func (Definition) TableName() string {
	return "definition"
}

type Example struct {
	DefinitionId int     `gorm:"column:definition_id"`
	Foreign      string  `gorm:"column:foreign"`
	Native       *string `gorm:"column:native"`
}

func (Example) TableName() string {
	return "example"
}

func AddWord(db *gorm.DB, word *Word) error {
	err := db.Create(word).Error
	return err
}

func GetWord(db *gorm.DB, id int) (Word, error) {
	word := Word{}
	if err := db.
		First(&word, id).Error; err != nil {
		return Word{}, err
	}

	return word, nil
}

func (w *Word) GetDefinitions(db *gorm.DB) ([]Definition, error) {
	defs := []Definition{}
	if err := db.
		Where("word_id = ?", w.Id).
		Find(&defs).Error; err != nil {
		return []Definition{}, err
	}
	return defs, nil
}

func (d *Definition) GetExamples(db *gorm.DB) ([]Example, error) {
	examples := []Example{}
	if err := db.
		Where("definition_id = ?", d.Id).
		Find(&examples).Error; err != nil {
		return []Example{}, err
	}

	return examples, nil
}
