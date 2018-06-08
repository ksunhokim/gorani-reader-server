package dbh

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type UnknownWord struct {
	UserId   int  `gorm:"column:user_id"`
	WordId   int  `gorm:"column:word_id"`
	Mastered bool `gorm:"column:unknown_word_mastered"`
}

func (UnknownWord) TableName() string {
	return "unknown_word"
}

type UnknownWordSource struct {
	UserId       int     `gorm:"column:user_id"`
	WordId       int     `gorm:"column:word_id" json:"word_id"`
	DefinitionId int     `gorm:"column:definition_id" json:"definition_id"`
	Book         *string `gorm:"column:unknown_word_source_book" json:"source_book"`
	Sentence     *string `gorm:"column:unknown_word_source_sentence" json:"source_sentence"`
}

func (UnknownWordSource) TableName() string {
	return "unknown_word_soucre"
}

func (u *User) AddUnknownWord(db *gorm.DB, source *UnknownWordSource) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	word := UnknownWord{
		UserId:   u.Id,
		WordId:   source.WordId,
		Mastered: false,
	}
	err = tx.Create(&word).Error
	if err != nil && !strings.Contains(err.Error(), "Duplicate") {
		return
	}

	source.UserId = u.Id
	return
}
