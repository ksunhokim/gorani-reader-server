package dbh

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

type Wordbook struct {
	Id       []byte    `gorm:"column:wordbook_uuid;primary_key"`
	UserId   int       `gorm:"column:user_id"`
	Name     string    `gorm:"column:wordbook_name"`
	SeenDate time.Time `gorm:"column:wordbook_seen_date"`
}

func (w *Wordbook) MarshalJSON() ([]byte, error) {
	var id string
	if w.Id == nil {
		id = ""
	} else {
		uid, err := uuid.FromBytes(w.Id)
		if err != nil {
			id = ""
		} else {
			id = uid.String()
		}
	}
	t := w.SeenDate.Format(time.RFC3339)
	jsonResult := fmt.Sprintf(`{"name":%q,"uuid":%q,"seen_date":%q}`, w.Name, id, t)
	return []byte(jsonResult), nil
}

func (w *Wordbook) UnmarshalJSON(b []byte) error {
	u := util.M{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		return err
	}

	if name, ok := u["name"].(string); ok {
		w.Name = name
	} else {
		return fmt.Errorf("JSON error")
	}

	if uid, ok := u["uuid"].(string); ok {
		id, err := uuid.Parse(uid)
		if err != nil {
			return err
		}
		bytes, err := id.MarshalBinary()
		if err != nil {
			return err
		}
		w.Id = bytes
	} else {
		return fmt.Errorf("JSON error")
	}

	if seendate, ok := u["seen_date"].(string); ok {
		d, err := time.Parse(time.RFC3339, seendate)
		if err != nil {
			return err
		}
		w.SeenDate = d
	} else {
		w.SeenDate = time.Time{}
	}
	return nil
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

func (u *User) GetWordbook(db *gorm.DB, id []byte) (Wordbook, error) {
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
