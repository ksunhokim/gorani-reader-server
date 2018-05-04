package model

import (
	"fmt"
	"time"

	"github.com/sunho/gorani-reader/pkg/dbs"
	"github.com/sunho/gorani-reader/pkg/util"

	"gopkg.in/mgo.v2/bson"
)

type (
	Wordbook struct {
		UpdatedAt time.Time       `bson:"updated_at"`
		UserId    bson.ObjectId   `bson:"user_id"`
		Index     int             `bson:"index"`
		Name      string          `bson:"name"`
		Entries   []WordbookEntry `bson:"entries"`
	}

	WordbookEntry struct {
		WordRef
		Star     bool   `json:"star"`
		Sentence string `json:"sentence"`
		Book     string `json:"book"`
	}
)

func vailidateNewEntires(entries []WordbookEntry) bool {
	if !util.IsDistinctSlice(entries) {
		return false
	}
	for _, entry := range entries {
		if !ValidateWord(entry.WordRef) {
			return false
		}
	}
	return true
}

func (book *Wordbook) PutEntries(entries []WordbookEntry) error {
	if !vailidateNewEntires(entries) {
		return fmt.Errorf("Invalid form")
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
				"entries":    entries,
				"updated_at": time.Now(),
			},
		})

	return err
}
