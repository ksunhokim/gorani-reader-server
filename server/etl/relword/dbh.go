package relword

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type relevantWord struct {
	WordId       int32  `gorm:"column:word_id"`
	TargetWordId int32  `gorm:"column:target_word_id"`
	RelType      string `gorm:"column:relevant_word_type"`
	Score        int32  `gorm:"column:relevant_word_score"`
	VoteSum      int32  `gorm:"column:relevant_word_vote_sum"`
}

func (relevantWord) TableName() string {
	return "relevant_word"
}

// normalized in order to keep vote data after renewing relevantWords
type relevantWordVote struct {
	WordId       int32  `gorm:"column:word_id"`
	TargetWordId int32  `gorm:"column:target_word_id"`
	UserId       int32  `gorm:"column:user_id"`
	RelType      string `gorm:"column:relevant_word_type"`
}

func (relevantWordVote) TableName() string {
	return "relevant_word_vote"
}

func deleteRelevantWords(db *gorm.DB, reltype string) error {
	err := db.
		Where("relevant_word_type = ?", reltype).
		Delete(&relevantWord{}).Error
	return err
}

// GOOORRRRRMM !@#!@#
func batchInsertRelevantWords(db *gorm.DB, words []relevantWord) error {
	placeholders := make([]string, 0, len(words))
	args := make([]interface{}, 0, len(words)*5)
	for _, word := range words {
		placeholders = append(placeholders, "(?, ?, ?, ?, ?)")
		args = append(args,
			word.WordId, word.TargetWordId, word.RelType,
			word.Score, word.VoteSum)
	}
	stmt := fmt.Sprintf(`INSERT INTO relevant_word 
		(word_id, target_word_id, relevant_word_type, 
		relevant_word_score, relevant_word_vote_sum) VALUES %s;`,
		strings.Join(placeholders, ","))
	err := db.Exec(stmt, args...).Error
	return err
}

func addRelevantWords(db *gorm.DB, reltype string, graph RelGraph) error {
	// maximun # of mysql placeholder is 65536
	// 5(# of columns of relevant_word) * 10000 = 50000 50000 < 65536
	bufferSize := 10000
	buffer := make([]relevantWord, 0, bufferSize)

	for _, v := range graph {
		for _, e := range v.Edges {
			word := relevantWord{
				WordId:       v.WordId,
				TargetWordId: e.TargetId,
				RelType:      reltype,
				Score:        e.Score,
				VoteSum:      0,
			}
			buffer = append(buffer, word)

			if len(buffer) == bufferSize {
				err := batchInsertRelevantWords(db, buffer)
				if err != nil {
					return err
				}

				buffer = make([]relevantWord, 0, bufferSize)
			}
		}
	}

	// flush remaining buffer
	if len(buffer) != 0 {
		err := batchInsertRelevantWords(db, buffer)
		if err != nil {
			return err
		}
	}

	return nil
}
