package models_test

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/gorani"
)

func Setup() *gorani.Gorani {
	bytes, err := ioutil.ReadFile("../config_test.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := config.NewConfig(bytes)
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.NewGorani(conf)
	if err != nil {
		panic(err)
	}
	setupDB(gorn.Mysql)

	return gorn
}

const (
	TestWordbookUuid = "3f06af63-a93c-11e4-9797-00505690773f"
	TestUserId       = 1
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

	db.Exec(fmt.Sprintf(`
	INSERT INTO wordbook 
		(wordbook_uuid, user_id, wordbook_is_unknown, wordbook_name, wordbook_update_date)
	VALUES
		(UNHEX(REPLACE('%s', '-', '')), 1, FALSE, 'test', NOW());`, TestWordbookUuid))

}
