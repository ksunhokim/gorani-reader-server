package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

type Wordbook struct {
	Id         util.UUID    `gorm:"column:wordbook_uuid;primary_key" json:"uuid"`
	Name       string       `gorm:"column:wordbook_name" json:"name"`
	SeenDate   util.RFCTime `gorm:"column:wordbook_seen_date" json:"seen_date"`
	UpdateDate util.RFCTime `gorm:"column:wordbook_update_date" json:"update_date"`
}

func (Wordbook) TableName() string {
	return "wordbook"
}

type UnknownWordbook struct {
	UserId     int32     `gorm:"column:user_id"`
	WordbookId util.UUID `gorm:"column:wordbook_uuid"`
}

func (UnknownWordbook) TableName() string {
	return "unknown_wordbook"
}

type UserWordbook struct {
	UserId     int32     `gorm:"column:user_id"`
	WordbookId util.UUID `gorm:"column:wordbook_uuid"`
}

func (UserWordbook) TableName() string {
	return "user_wordbook"
}

func (wb *Wordbook) Update(db *gorm.DB) error {
	err := db.Save(wb).Error
	return err
}

func (wb *Wordbook) Delete(db *gorm.DB) error {
	err := db.Delete(&wb).Error
	return err
}

func getWordbook(db *gorm.DB, id util.UUID) (Wordbook, error) {
	wordbook := Wordbook{}
	if err := db.
		Where("wordbook_uuid = ?", id).
		First(&wordbook).
		Error; err != nil {
		return Wordbook{}, err
	}

	return wordbook, nil
}

func (u *User) GetWordbook(db *gorm.DB, id util.UUID) (Wordbook, error) {
	userwordbook := UserWordbook{}
	if err := db.
		Where("wordbook_uuid = ? AND user_id = ?", id, u.Id).
		First(&userwordbook).
		Error; err != nil {
		return Wordbook{}, err
	}

	wordbook, err := getWordbook(db, userwordbook.WordbookId)
	return wordbook, err
}

func (u *User) GetWordbooks(db *gorm.DB) ([]Wordbook, error) {
	wordbooks := []Wordbook{}
	if err := db.
		Raw(`SELECT 
				wb.*
			FROM
				wordbook wb
			INNER JOIN
				user_wordbook uw
			ON
				uw.wordbook_uuid  = wb.wordbook_uuid
			WHERE
				uw.user_id = ?;`, u.Id).
		Scan(&wordbooks).Error; err != nil {
		return nil, err
	}

	return wordbooks, nil
}

func (u *User) AddWordbook(db *gorm.DB, wordbook *Wordbook) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	wordbook.UpdateDate = util.RFCTime{t}

	err = tx.Create(wordbook).Error
	if err != nil {
		return
	}

	userwordbook := UserWordbook{
		UserId:     u.Id,
		WordbookId: wordbook.Id,
	}
	err = tx.Create(&userwordbook).Error
	return
}

func (u *User) GetUnknownWordbook(db *gorm.DB) (Wordbook, error) {
	uwordbook := UnknownWordbook{}
	if err := db.
		Where("user_id = ?", u.Id).
		First(&uwordbook).
		Error; err != nil {
		return Wordbook{}, err
	}

	wordbook, err := getWordbook(db, uwordbook.WordbookId)
	return wordbook, err
}
