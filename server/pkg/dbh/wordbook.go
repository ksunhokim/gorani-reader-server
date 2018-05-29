package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Wordbook struct {
	Id       []byte    `gorm:"column:wordbook_uuid;primary_key"`
	UserId   int       `gorm:"column:user_id"`
	Name     string    `gorm:"column:wordbook_name"`
	SeenDate time.Time `gorm:"column:wordbook_seen_date"`
}

func (Wordbook) TableName() string {
	return "wordbook"
}

func (wb *Wordbook) Update(db *gorm.DB) error {
	err := db.Save(wb).Error
	return err
}

func (wb *Wordbook) Delete(db *gorm.DB) error {
	err := db.Delete(wb).Error
	return err
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
		Where("user_id = ?", u.Id).
		Find(&wordbooks).
		Error; err != nil {
		return []Wordbook{}, err
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

	wordbook.SeenDate = time.Now().UTC()
	wordbook.UserId = u.Id
	if err = tx.Create(wordbook).Error; err != nil {
		return err
	}

	t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	date := WordbookEntriesUpdateDate{
		WordbookId: wordbook.Id,
		Date:       t,
	}
	err = tx.Create(&date).Error

	return err
}
