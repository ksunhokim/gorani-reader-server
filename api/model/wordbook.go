package model

import (
	"fmt"
	"time"

	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/util"

	"gopkg.in/mgo.v2/bson"
)

type (
	Wordbook struct {
		UpdatedAt *time.Time      `bson:"updated_at"`
		UserId    bson.ObjectId   `bson:"user_id"`
		Index     int             `json:"index" bson:"index"`
		Name      string          `json:"name" bson:"name"`
		Entries   []WordbookEntry `json:"entries" bson:"entries"`
	}

	WordbookEntry struct {
		WordRef
		Star     bool   `json:"star"`
		Sentence string `json:"sentence"`
		Book     string `json:"book"`
	}

	WordRef struct {
		Word       string `json:"word"`
		Definition uint   `json:"definition"`
	}
)

func (book *Wordbook) PutEntries(entries []WordbookEntry) error {
	if !util.IsDistinctSlice(entries) {
		return fmt.Errorf("No distinct entries")
	}
	for _, entry := range entries {
		if !ValidateWord(entry.WordRef) {
			return fmt.Errorf("No such word")
		}
	}

	sess := dbs.MDB.Copy()
	defer sess.Close()
	err := sess.DB("").C("wordbooks").Update(
		bson.M{
			"user_id": book.UserId,
			"index":   book.Index,
		},
		bson.M{
			"$set": bson.M{
				"entries": entries,
			},
		})

	return err
}
