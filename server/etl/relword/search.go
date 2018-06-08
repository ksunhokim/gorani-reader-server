package relword

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

func FindKnowns(db *gorm.DB, reltype string, word dbh.Word, user dbh.User, maxresult int) (words []dbh.Word, err error) {
	err = db.Raw(`
		SELECT word.* 
		FROM
			relevant_word rw 
		INNER JOIN
			known_word nw
		ON
			rw.target_word_id = nw.word_id
		INNER JOIN
			word
		ON
			word.word_id = rw.word_id
		WHERE
			rw.relevant_word_type = ? AND
			rw.word_id = ? AND
			nw.user_id = ?
		ORDER BY
			rw.relevant_word_score DESC,
			rw.relevant_word_vote_sum DESC
		LIMIT ?;`, reltype, word.Id, user.Id, maxresult).
		Scan(&words).Error
	return
}
