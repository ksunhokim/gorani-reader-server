package dbh

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	Id            int32        `gorm:"column:word_id;primary_key" json:"id"`
	Word          string       `gorm:"column:word;not null;unique" json:"word"`
	Pronunciation *string      `gorm:"column:word_pronunciation" json:"pronunciation,omitempty"`
	Definitions   []Definition `json:"definitions,omitempty"`
}

func (Word) TableName() string {
	return "word"
}

type Definition struct {
	Id         int32     `gorm:"column:definition_id;primary_key" json:"id"`
	WordId     int32     `gorm:"column:word_id;not null" json:"word_id"`
	Definition string    `gorm:"column:definition;not null" json:"definition"`
	POS        *string   `gorm:"column:definition_pos" json:"pos,omitempty"`
	Examples   []Example `json:"examples,omitempty"`
}

func (Definition) TableName() string {
	return "definition"
}

type Example struct {
	DefinitionId int32   `gorm:"column:definition_id;not null" json:"definition_id"`
	Foreign      string  `gorm:"column:foreign;not null" json:"foreign"`
	Native       *string `gorm:"column:native" json:"native,omitempty"`
}

func (Example) TableName() string {
	return "example"
}

func AddWord(db *gorm.DB, word *Word) error {
	err := db.Create(word).Error
	return err
}

func GetWordById(db *gorm.DB, id int32) (Word, error) {
	word := Word{}
	if err := db.
		Preload("Definitions").
		Preload("Definitions.Examples").
		First(&word, id).Error; err != nil {
		return Word{}, err
	}

	return word, nil
}

func GetWordByWord(db *gorm.DB, str string) (Word, error) {
	word := Word{}
	if err := db.
		Preload("Definitions").
		Preload("Definitions.Examples").
		Where("word = ?", str).
		First(&word).Error; err != nil {
		return Word{}, err
	}

	return word, nil
}

func GetWords(db *gorm.DB) ([]Word, error) {
	words := []Word{}
	if err := db.
		Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (w *Word) Delete(db *gorm.DB) error {
	err := db.Delete(w).Error
	return err
}
