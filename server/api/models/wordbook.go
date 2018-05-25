package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Wordbook struct {
	Id         []byte    `gorm:"column:wordbook_uuid;primary_key"`
	UserId     int       `gorm:"column:user_id"`
	IsUnknown  bool      `gorm:"column:wordbook_is_unknown"`
	Name       string    `gorm:"column:wordbook_name"`
	UpdateDate time.Time `gorm:"column:wordbook_update_date"`
}

func (Wordbook) TableName() string {
	return "wordbook"
}

func GetWordbook(db *gorm.DB, id []byte) (Wordbook, error) {
	wordbook := Wordbook{}
	if err := db.
		Where("wordbook_uuid = ?", id).
		First(&wordbook).
		Error; err != nil {
		return Wordbook{}, err
	}
	return wordbook, nil
}

func (u *User) GetWordbooks(db *gorm.DB) ([]Wordbook, error) {
	wordbooks := []Wordbook{}
	if err := db.
		Where("user_id = ? AND wordbook_is_unknown = FALSE", u.Id).
		Find(&wordbooks).
		Error; err != nil {
		return []Wordbook{}, err
	}
	return wordbooks, nil
}

func (u *User) CreateWordbook(db *gorm.DB, wordbook *Wordbook) error {
	if wordbook.IsUnknown {
		return fmt.Errorf("Unknown wordbook cannot be created manually")
	}

	wordbook.UpdateDate = time.Now().UTC()
	wordbook.UserId = u.Id
	err := db.Create(wordbook).Error

	return err
}
