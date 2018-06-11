package sentencer

import (
	"io"
	"sync"
)

type instance struct {
	mu sync.RWMutex
	s  *Sentencer
}

func (i *instance) get() *Sentencer {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.s
}

func (i *instance) set(s *Sentencer) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.s = s
}

var ins instance

func init() {
	s := New(
		make(Dictionary),
		DotSpecialCases{},
		NewStemmer(
			make(map[string]string),
			make(map[string]string),
		),
	)

	ins = instance{s: s}
}

func SetSentencer(s *Sentencer) {
	ins.set(s)
}

func ExtractSentencesFromText(r io.Reader) []Sentence {
	return ins.get().ExtractSentencesFromText(r)
}

func ExtractSentencesFromHtml(r io.Reader) ([]Sentence, error) {
	return ins.get().ExtractSentencesFromHtml(r)
}
