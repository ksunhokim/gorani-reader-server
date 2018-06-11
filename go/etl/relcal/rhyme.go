package relcal

import (
	"errors"
	"strings"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

type wordSet map[int]struct{}

func (a wordSet) Sub(b wordSet) wordSet {
	out := make(wordSet, len(a))
	for i := range a {
		if _, ok := b[i]; !ok {
			out[i] = struct{}{}
		}
	}
	return out
}

func (a wordSet) Intersect(b wordSet) wordSet {
	out := make(wordSet, len(a))
	for i := range a {
		if _, ok := b[i]; ok {
			out[i] = struct{}{}
		}
	}
	return out
}

type rhymeCalculator struct {
}

const maxToken = 39

var arpMap = map[string]int{"AA": 0, "AE": 1, "AH": 2, "AO": 3, "AW": 4, "AY": 5, "B": 6, "CH": 7, "D": 8, "DH": 9, "EH": 10, "ER": 11, "EY": 12, "F": 13, "G": 14, "HH": 15, "IH": 16, "IY": 17, "JH": 18, "K": 19, "L": 20, "M": 21, "N": 22, "NG": 23, "OW": 24, "OY": 25, "P": 26, "R": 27, "S": 28, "SH": 29, "T": 30, "TH": 31, "UH": 32, "UW": 33, "V": 34, "W": 35, "Y": 36, "Z": 37, "ZH": 38}

var (
	ErrUnknownToken     = errors.New("relcal: Unknown token in word's pronunciation")
	ErrNilPronunciaions = errors.New("relcal: Every word's pronunciation is nil")
)

func init() {
	err := calculators.add(rhymeCalculator{})
	if err != nil {
		panic(err)
	}
}

func (rhymeCalculator) Calculate(words []dbh.Word, minscore int) (graph Graph, err error) {
	// [token's offset in the syllables counting from the back][token number]wordSet
	sufTokWordArr := [][maxToken]wordSet{}

	// create sufTokWordArr
	for _, word := range words {
		if word.Pronunciation == nil {
			continue
		}
		pron := *word.Pronunciation

		rawToks := strings.Split(pron, " ")
		for i := range rawToks {
			// reverse direction
			val := rawToks[len(rawToks)-i-1]

			// this will not allow empty string
			tok, ok := arpMap[val]
			if !ok {
				err = ErrUnknownToken
				return
			}

			// expand the array
			if len(sufTokWordArr) == i {
				new := [maxToken]wordSet{}
				for i := 0; i < maxToken; i++ {
					// if lack of memory, reduce the cap
					new[i] = make(wordSet, 1000)
				}
				sufTokWordArr = append(sufTokWordArr, new)
			}

			sufTokWordArr[i][tok][word.Id] = struct{}{}
		}
	}

	if len(sufTokWordArr) == 0 {
		err = ErrNilPronunciaions
		return
	}

	// count score
	for _, word := range words {
		if word.Pronunciation == nil {
			continue
		}
		pron := *word.Pronunciation
		rawToks := strings.Split(pron, " ")

		vertex := Vertex{
			WordId: word.Id,
			Edges:  []Edge{},
		}

		val := rawToks[len(rawToks)-1]
		tok := arpMap[val]
		// set of words that also have the same last token
		set := sufTokWordArr[0][tok]

		for i := 1; i < len(rawToks); i++ {
			val := rawToks[len(rawToks)-i-1]
			tok := arpMap[val]

			// set of words that have the same i+1 or more tokens
			ins := set.Intersect(sufTokWordArr[i][tok])

			// set of words that have the same i tokens
			rem := set.Sub(ins)
			set = ins

			for ele := range rem {
				score := i
				if minscore > score {
					continue
				}
				edge := Edge{
					TargetId: ele,
					Score:    int(score),
				}
				vertex.Edges = append(vertex.Edges, edge)
			}

			if len(set) == 0 {
				break
			}
		}

		// this happens when the loop above terminated by for statement's condition
		if len(set) != 0 {
			for ele := range set {
				if ele == vertex.WordId {
					continue
				}

				score := len(rawToks)
				if minscore > score {
					continue
				}
				edge := Edge{
					TargetId: ele,
					Score:    int(score),
				}
				vertex.Edges = append(vertex.Edges, edge)
			}
		}

		graph.Vertexs = append(graph.Vertexs, vertex)
	}

	return
}

func (rhymeCalculator) RelType() string {
	return "rhyme"
}
