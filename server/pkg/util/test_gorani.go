package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

func SetupTestGorani() *gorani.Gorani {
	conf, err := gorani.NewConfig("../../config_test.yaml")
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.New(conf)
	if err != nil {
		panic(err)
	}
	setupDB(gorn.Mysql)

	return gorn
}

const (
	TestUserId = 1
)

func setupDB(db *gorm.DB) {
	rows, err := db.DB().Query(`
	SELECT CONCAT('delete from ', table_name, ';') FROM 
		information_schema.tables 
	WHERE table_schema=(SELECT DATABASE()) AND
	table_type = 'BASE TABLE';`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		query := ""
		rows.Scan(&query)
		if err := db.Exec(query).Error; err != nil {
			panic(err)
		}
	}

	db.Exec(fmt.Sprintf(`
	INSERT INTO user 
		(user_id, user_name)
	VALUES
		(%d, 'test');`, TestUserId))

	db.Exec(`
	INSERT INTO oauth_passport
		(user_id, oauth_service, oauth_user_id)
	VALUES
		(1, 'hoi', 'asdf');`)

	db.Exec(`
	INSERT INTO word
		(word_id, word, word_pronunciation)
	VALUES
		(1, 'test', NULL);`)

	db.Exec(`
	INSERT INTO word
		(word_id, word, word_pronunciation)
	VALUES
		(2, 'test2', NULL);`)

	db.Exec(`
	INSERT INTO definition
		(definition_id, word_id, definition_pos, definition)
	VALUES
		(1, 1, NULL, 'test');`)

	db.Exec(`
	INSERT INTO definition
		(definition_id, word_id, definition_pos, definition)
	VALUES
		(2, 1, NULL, 'test');`)

}
