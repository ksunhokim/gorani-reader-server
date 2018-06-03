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

func (s *Sentencer) ExtractSentencesFromText(str string) (out []Sentence) {
	t := s.createTokenizer(strings.NewReader(str))
	toks := t.Tokenize()
	i := 0
	out = append(out, Sentence{Words: []WordId{}})

	for _, tok := range toks {
		out[i].Origin += tok.Lit
		if isWordToken(tok) {
			word := strings.ToLower(tok.Lit)
			if id, ok := s.Dict[word]; ok {
				out[i].Words = append(out[i].Words, id)
				continue
			}
			word = s.Stemmer.Stem(word)
			if id, ok := s.Dict[word]; ok {
				out[i].Words = append(out[i].Words, id)
				continue
			}
		}
		if tok.Kind == TokenKindEos {
			i++
			out = append(out, Sentence{Words: []WordId{}})
		}
	}
	return
}

func (s *Sentencer) ExtractSentencesFromHtml(raw string) (sens []Sentence, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		return
	}

	doc.Find("p").Each(func(i int, sel *goquery.Selection) {
		str := sel.Text()
		sens = append(sens, s.ExtractSentencesFromText(str)...)
	})

	return
}
