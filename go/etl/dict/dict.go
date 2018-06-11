package dict

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
)

func Create(mysql *gorm.DB) (string, error) {
	url := fmt.Sprintf("/tmp/dict_%d.db", time.Now().UnixNano())
	db, err := gorm.Open("sqlite3", url)
	db.LogMode(true)
	if err != nil {
		return "", err
	}
	defer db.Close()

	err = migrate(db)
	if err != nil {
		return "", err
	}

	words, err := dbh.GetWords(mysql)
	if err != nil {
		return "", err
	}

	// words.definitions are not fully loaded yet
	for i := range words {
		fullWord, err := dbh.GetWordById(mysql, words[i].Id)
		if err != nil {
			return "", err
		}

		words[i] = fullWord
	}

	err = insert(db, words)
	if err != nil {
		return "", err
	}

	return url, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&dbh.Word{}, &dbh.Definition{}, &dbh.Example{}).Error
	if err != nil {
		return err
	}

	db.Model(&dbh.Definition{}).AddIndex("idx_definition", "word_id")
	db.Model(&dbh.Example{}).AddIndex("idx_example", "definition_id")
	return err
}

func insert(db *gorm.DB, words []dbh.Word) error {
	db = db.Set("gorm:save_associations", false)

	for _, word := range words {
		if err := db.
			Create(&word).Error; err != nil {
			return err
		}

		defs := word.Definitions

		for _, def := range defs {
			if err := db.
				Create(&def).Error; err != nil {
				return err
			}
			examples := def.Examples

			for _, example := range examples {
				if err := db.
					Create(&example).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
