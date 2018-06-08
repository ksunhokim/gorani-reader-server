package dbh

import (
	"github.com/jinzhu/gorm"
)

type RelevantWord struct {
	WordId       int    `gorm:"column:word_id"`
	TargetWordId int    `gorm:"column:target_word_id"`
	RelType      string `gorm:"column:relevant_word_type"`
	Score        int    `gorm:"column:relevant_word_score"`
	VoteSum      int    `gorm:"column:relevant_word_vote_sum"`
}

func (RelevantWord) TableName() string {
	return "relevant_word"
}

// normalized in order to keep vote data after renewing relevantWords
type RelevantWordVote struct {
	WordId       int    `gorm:"column:word_id"`
	TargetWordId int    `gorm:"column:target_word_id"`
	UserId       int    `gorm:"column:user_id"`
	RelType      string `gorm:"column:relevant_word_type"`
}

func (RelevantWordVote) TableName() string {
	return "relevant_word_vote"
}

// c should be closed manually
func StreamAddRelevantWords(db *gorm.DB, c chan RelevantWord) <-chan error {
	c2 := make(chan []interface{})
	errC := streamAddRows(db, "relevant_word",
		[]string{"word_id", "target_word_id", "relevant_word_type",
			"relevant_word_score", "relevant_word_vote_sum"}, c2)

	go func() {
		for {
			select {
			case word, more := <-c:
				if !more {
					close(c2)
					return
				}
				c2 <- []interface{}{word.WordId, word.TargetWordId, word.RelType,
					word.Score, word.VoteSum}
			}
		}
	}()

	return errC
}

func (u *User) FindRelevantKnownWords(db *gorm.DB, reltype string, word Word, maxresult int) (words []Word, err error) {
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
		LIMIT ?;`, reltype, word.Id, u.Id, maxresult).
		Scan(&words).Error
	return
}

func DeleteRelevantWords(db *gorm.DB, reltype string) error {
	err := db.
		Where("relevant_word_type = ?", reltype).
		Delete(&RelevantWord{}).Error
	return err
}
