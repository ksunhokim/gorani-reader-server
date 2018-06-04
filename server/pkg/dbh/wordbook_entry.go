package dbh

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type WordbookEntry struct {
	WordbookId     UUID    `gorm:"column:wordbook_uuid;primary_key" json:"-"`
	DefinitionId   int     `gorm:"column:definition_id" json:"definition_id"`
	SourceBook     *string `gorm:"column:wordbook_entry_source_book" json:"source_book;omitempty"`
	SourceSentence *string `gorm:"column:wordbook_entry_source_sentence" json:"source_sentence;omitempty"`
	WordIndex      *int    `gorm:"column:wordbook_entry_word_index" json:"word_inder;omitempty"`
	AddedDate      RFCTime `gorm:"column:wordbook_entry_added_date" json:"added_date"`
}

func (WordbookEntry) TableName() string {
	return "wordbook_entry"
}

func (wb *Wordbook) GetEntries(db *gorm.DB) ([]WordbookEntry, error) {
	out := []WordbookEntry{}
	if err := db.
		Where(`wordbook_uuid = ?`, wb.Id).
		Find(&out).Error; err != nil {
		return []WordbookEntry{}, err
	}
	return out, nil
}

func (wb *Wordbook) ReloadLockInShareMode(tx *gorm.DB) error {
	if err := tx.
		Raw(`SELECT
			* 
		FROM 
		wordbook
		WHERE
			wordbook_uuid = ?
		LOCK IN SHARE MODE;`,
			wb.Id).
		Scan(wb).Error; err != nil {
		return err
	}
	return nil
}

func (wb *Wordbook) AddEntry(db *gorm.DB, date time.Time, entry *WordbookEntry) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = wb.ReloadLockInShareMode(tx)
	if err != nil {
		return
	}

	if wb.UpdateDate.After(date) {
		return fmt.Errorf("Trying to use old value")
	}

	entry.WordbookId = wb.Id
	if err = tx.
		Create(entry).Error; err != nil {
		return
	}

	wb.UpdateDate = RFCTime{date}
	err = wb.Update(tx)
	return
}

func (wb *Wordbook) UpdateEntries(db *gorm.DB, date time.Time, entries []WordbookEntry) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	err = wb.ReloadLockInShareMode(tx)
	if err != nil {
		return
	}

	if wb.UpdateDate.After(date) {
		return fmt.Errorf("Trying to use old value")
	}

	if err = tx.
		Where("wordbook_uuid = ?", wb.Id).
		Delete(WordbookEntry{}).Error; err != nil {
		return
	}

	for _, entry := range entries {
		entry.WordbookId = wb.Id
		if err = tx.
			Create(&entry).
			Error; err != nil {
			return
		}
	}

	wb.UpdateDate = RFCTime{date}
	err = wb.Update(tx)
	return
}
