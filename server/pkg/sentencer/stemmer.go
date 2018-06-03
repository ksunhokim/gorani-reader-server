package sentencer

import "strings"

type Stemmer struct {
	IrPast     map[string]string
	IrComplete map[string]string
}

func NewStemmer(irregularPast map[string]string, irregularComplete map[string]string) *Stemmer {
	return &Stemmer{
		IrPast:     irregularPast,
		IrComplete: irregularComplete,
	}
}

func (s *Stemmer) Stem(word string) string {
	if w, ok := s.IrComplete[word]; ok {
		return w
	}

	if w, ok := s.IrPast[word]; ok {
		return w
	}

	if strings.HasSuffix(word, "es") {
		return word[:len(word)-2]
	}

	if strings.HasSuffix(word, "s") {
		return word[:len(word)-1]
	}

	if strings.HasSuffix(word, "ied") {
		return word[:len(word)-3] + "y"
	}

	if strings.HasSuffix(word, "ed") {
		if len(word) >= 4 && word[len(word)-3] == word[len(word)-3] {
			return word[:len(word)-3]
		}
		return word[:len(word)-2]
	}

	if strings.HasSuffix(word, "ying") {
		return word[:len(word)-4] + "ie"
	}

	if strings.HasSuffix(word, "ing") {
		if len(word) >= 5 && word[len(word)-4] == word[len(word)-5] {
			return word[:len(word)-4]
		}
		return word[:len(word)-3]
	}

	return word
}
