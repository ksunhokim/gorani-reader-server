package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

type Wordbook struct {
	Id         util.UUID    `gorm:"column:wordbook_uuid;primary_key" json:"uuid"`
	UserId     int          `gorm:"column:user_id" json:"-"`
	Name       string       `gorm:"column:wordbook_name" json:"name"`
	SeenDate   util.RFCTime `gorm:"column:wordbook_seen_date" json:"seen_date"`
	UpdateDate util.RFCTime `gorm:"column:wordbook_update_date" json:"update_date"`
}

func (Wordbook) TableName() string {
	return "wordbook"
}

func (wb *Wordbook) Update(db *gorm.DB) error {
	err := db.Save(wb).Error
	return err
}

func (wb *Wordbook) Delete(db *gorm.DB) error {
	err := db.Delete(&wb).Error
	return err
}

func (u *User) GetWordbook(db *gorm.DB, id util.UUID) (Wordbook, error) {
	wordbook := Wordbook{}
	if err := db.
		Where("wordbook_uuid = ? AND user_id = ?", id, u.Id).
		First(&wordbook).
		Error; err != nil {
		return Wordbook{}, err
	}

	return wordbook, nil
}

func (u *User) GetWordbooks(db *gorm.DB) ([]Wordbook, error) {
	wordbooks := []Wordbook{}
	if err := db.
		Where("user_id = ?", u.Id).
		Find(&wordbooks).
		Error; err != nil {
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

	wordbook.UserId = u.Id
	t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	wordbook.UpdateDate = util.RFCTime{t}

	err = tx.Create(wordbook).Error
	return err
}
