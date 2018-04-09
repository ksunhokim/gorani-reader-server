package model

import (
	"github.com/go-bongo/bongo"
)

type (
	Wordbook struct {
		bongo.DocumentBase `bson:",inline" json:"-"`
		Name               string          `json:"name"`
		Entries            []WordbookEntry `json:"entries"`
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

// may bottleneck
func (book *Wordbook) ValidateWord(ref WordRef) bool {
	if !ValidateWord(ref) {
		return false
	}
	for _, entry := range book.Entries {
		if entry.Word == ref.Word && entry.Definition == ref.Definition {
			return false
		}
	}
	return true
}
