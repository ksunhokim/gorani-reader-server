package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type WordbookEntry struct {
	WordbookId     []byte    `gorm:"column:wordbook_uuid;primary_key"`
	DefinitionId   int       `gorm:"column:definition_id"`
	SourceBook     *string   `gorm:"column:wordbook_entry_source_book"`
	SourceSentence *string   `gorm:"column:wordbook_entry_source_sentence"`
	WordIndex      *int      `gorm:"column:wordbook_entry_word_index"`
	AddedDate      time.Time `gorm:"column:wordbook_entry_added_date"`
}

func (WordbookEntry) TableName() string {
	return "wordbook_entry"
}

type WordbookEntriesUpdateDate struct {
	WordbookId []byte    `gorm:"column:wordbook_uuid;primary_key"`
	Date       time.Time `gorm:"column:wordbook_entry_update_date"`
}

func (WordbookEntriesUpdateDate) TableName() string {
	return "wordbook_entries_update_date"
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

func (wb *Wordbook) getEntriesUpdateDate(db *gorm.DB) (WordbookEntriesUpdateDate, error) {
	date := WordbookEntriesUpdateDate{}
	if err := db.
		Raw(`SELECT
				* 
			FROM 
			wordbook_entries_update_date 
			WHERE
				wordbook_uuid = ?
			LOCK IN SHARE MODE;`,
			wb.Id).
		Scan(&date).Error; err != nil {
		return WordbookEntriesUpdateDate{}, err
	}
	return date, nil
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

	date2, err := wb.getEntriesUpdateDate(tx)
	if err != nil {
		return err
	}

	if date2.Date.After(date) {
		return fmt.Errorf("Trying to use old value")
	}

	entry.WordbookId = wb.Id
	if err = tx.
		Create(entry).Error; err != nil {
		return err
	}

	date2.Date = date
	if err = tx.
		Save(&date2).Error; err != nil {
		return err
	}

	return nil
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

	date2, err := wb.getEntriesUpdateDate(tx)
	if err != nil {
		return err
	}

	if date2.Date.After(date) {
		return fmt.Errorf("Trying to use old value")
	}

	if err = tx.
		Where("wordbook_uuid = ?", wb.Id).
		Delete(WordbookEntry{}).Error; err != nil {
		return err
	}

	for _, entry := range entries {
		entry.WordbookId = wb.Id
		if err = tx.
			Create(&entry).
			Error; err != nil {
			return err
		}
	}

	date2.Date = date
	if err = tx.
		Save(&date2).Error; err != nil {
		return err
	}

	return nil
}
