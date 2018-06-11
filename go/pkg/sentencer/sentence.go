package sentencer

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Sentencer struct {
	Dict            Dictionary
	DotSpecialCases DotSpecialCases
	Stemmer         *Stemmer
}

func New(dict Dictionary, dot DotSpecialCases, stemmer *Stemmer) *Sentencer {
	return &Sentencer{
		Dict:            dict,
		DotSpecialCases: dot,
		Stemmer:         stemmer,
	}
}

func (s *Sentencer) createTokenizer(r io.Reader) *Tokenizer {
	t := NewTokenizer(r)
	t.DotSpecialCases = s.DotSpecialCases
	return t
}

func (s *Sentencer) ExtractSentencesFromText(r io.Reader) (out []Sentence) {
	t := s.createTokenizer(r)
	toks := t.Tokenize()

	i := 0
	out = append(out, Sentence{Words: []WordId{}})
	for _, tok := range toks {
		out[i].Origin += tok.Lit

		if isWordToken(tok) {
			// also add raw word if it exists in dictionary
			word := strings.ToLower(tok.Lit)
			id, ok := s.Dict[word]
			if ok {
				out[i].Words = append(out[i].Words, id)
			}

			word = s.Stemmer.Stem(word)
			sid, ok := s.Dict[word]
			if ok && sid != id {
				out[i].Words = append(out[i].Words, sid)
			}
		}

		if tok.Kind == TokenKindEos {
			i++
			out = append(out, Sentence{Words: []WordId{}})
		}
	}

	return
}

func (s *Sentencer) ExtractSentencesFromHtml(r io.Reader) (sens []Sentence, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return
	}

	// every texts were in p tag as I analyzed some samples
	doc.Find("p").Each(func(i int, sel *goquery.Selection) {
		str := sel.Text()
		r := strings.NewReader(str)
		sens = append(sens, s.ExtractSentencesFromText(r)...)
	})

	return
}
