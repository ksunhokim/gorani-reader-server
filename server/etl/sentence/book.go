package sentence

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

func FindFromBook(db *gorm.DB, word1 dbh.Word, word2 dbh.Word, maxdistance int) (sentences []dbh.Sentence, err error) {
	err = db.Raw(`
		SELECT sentence.*
		FROM
			word_sentence a
		INNER JOIN
			word_sentence b
		ON 
			a.word_id = ? AND
			b.word_id = ? AND
			a.sentence_id = b.sentence_id AND
			ABS(a.word_position - b.word_position) <= ? 
		INNER JOIN
			sentence
		ON
			a.sentence_id = sentence.sentence_id
		GROUP BY
			sentence_id;`, word1.Id, word2.Id, maxdistance).
		Scan(&sentences).Error
	return
}
