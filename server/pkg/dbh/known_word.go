package dbh

import (
	"time"

	"github.com/jinzhu/gorm"
)

type KnownWord struct {
	UserId    int32     `gorm:"column:user_id"`
	WordId    int32     `gorm:"column:word_id"`
	AddedDate time.Time `gorm:"column:known_word_added_date"`
}

func (KnownWord) TableName() string {
	return "known_word"
}

func (u *User) AddKnownWord(db *gorm.DB, wordId int32) error {
	word := KnownWord{
		UserId:    u.Id,
		WordId:    wordId,
		AddedDate: time.Now().UTC(),
	}
	err := db.Create(&word).Error
	return err
}

func (u *User) GetKnownWords(db *gorm.DB) ([]int32, error) {
	words := []KnownWord{}
	if err := db.
		Where("user_id = ?", u.Id).
		Find(&words).Error; err != nil {
		return nil, err
	}

	arr := make([]int32, 0, len(words))
	for _, w := range words {
		arr = append(arr, w.WordId)
	}
	return arr, nil
}
