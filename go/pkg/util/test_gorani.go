package util

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader-server/go/pkg/gorani"
)

func SetupTestGorani() *gorani.Gorani {
	isCompose := os.Getenv("COMPOSE")

	conf := gorani.Config{}
	if isCompose == "true" {
		conf.RedisURL = "redis://redis/"
		conf.MysqlURL = "gorani:gorani@tcp(mysql:3306)/gorani_test?parseTime=true"
		conf.S3EndPoint = "s3:9000"
		conf.S3Id = "test1234"
		conf.S3Secret = "test1234"
		conf.S3Ssl = false
		conf.GoMaxProcs = 2
		conf.Debug = false
	} else {
		conf.RedisURL = "redis://127.0.0.1/"
		conf.MysqlURL = "gorani:gorani@tcp(127.0.0.1:3306)/gorani_test?parseTime=true"
		conf.S3EndPoint = "127.0.0.1:9000"
		conf.S3Id = "test1234"
		conf.S3Secret = "test1234"
		conf.S3Ssl = false
		conf.GoMaxProcs = 2
		conf.Debug = false
	}

	gorn, err := gorani.New(conf)
	if err != nil {
		panic(err)
	}
	setupDB(gorn.Mysql)

	return gorn
}

const (
	TestUserId   = 1
	TestBookIsbn = "asdf"
)

func mustExec(db *gorm.DB, str string) {
	err := db.Exec(str).Error
	if err != nil {
		panic(err)
	}
}

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

	mustExec(db, fmt.Sprintf(`
		INSERT INTO user 
			(user_id, user_name)
		VALUES
			(%d, 'test');`, TestUserId))

	mustExec(db, `
		INSERT INTO word
			(word_id, word, word_pronunciation)
		VALUES
			(1, 'test', NULL);`)

	mustExec(db, `
		INSERT INTO word
			(word_id, word, word_pronunciation)
		VALUES
			(2, 'test2', NULL);`)

	mustExec(db, `
		INSERT INTO definition
			(definition_id, word_id, definition_pos, definition)
		VALUES
			(1, 1, NULL, 'test');`)

	mustExec(db, `
		INSERT INTO definition
			(definition_id, word_id, definition_pos, definition)
		VALUES
			(2, 1, NULL, 'test');`)

	mustExec(db, `
		INSERT INTO book
			(book_isbn, book_name, book_author, book_cover_image)
		VALUES
			('asdf', 'asdf', 'asdf', 'asdf');`)

	mustExec(db, `
		INSERT INTO sentence
			(sentence_id, sentence, book_isbn)
		VALUES
			(1, 'test test2', 'asdf');`)

	mustExec(db, `
		INSERT INTO word_sentence
			(sentence_id, word_id, word_position)
		VALUES
			(1, 1, 0);`)

	mustExec(db, `
		INSERT INTO word_sentence
			(sentence_id, word_id, word_position)
		VALUES
			(1, 2, 1);`)

	mustExec(db, `
		INSERT INTO relevant_word_type
			(relevant_word_type_code, relevant_word_type_name)
		VALUES
			(1, 'test');`)

	mustExec(db, `
		INSERT INTO relevant_word
			(word_id, target_word_id, relevant_word_type_code, relevant_word_score, relevant_word_vote_sum)
		VALUES
			(1, 2, 1, 10, 1);`)

	mustExec(db, `
		INSERT INTO genre
			(genre_code, genre_name)
		VALUES
			(1, 'test');`)
}
