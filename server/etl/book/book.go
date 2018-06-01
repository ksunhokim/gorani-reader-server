package book

import (
	"io"

	"github.com/meskio/epubgo"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

type Book struct {
	Title       string
	Author      string
	Genre       string
	CoverExt    string
	CoverReader io.Reader
	Sentences   []Sentence
	Reviews     []Review
}

type Review struct {
	Provider string
	Number   int
	Rate     float32
}

type Sentence []int

func Parse(dict Dictionary, r io.ReaderAt, size int64) (*Book, error) {
	file, err := epubgo.Load(r, size)
	if err != nil {
		return nil, err
	}

	iter, err := file.Spine()
	if err != nil {
		return nil, err
	}

	sentences, err := parseSentencesByIter(dict, iter)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func parseSentencesByIter(dict Dictionary, iter *epubgo.SpineIterator) ([]Sentence, error) {
	ri, err := iter.Open()
	if err != nil {
		return nil, err
	}
	sentences := []Sentence{}
	sentences2, err := parseSentences(dict, ri)
	if err == nil {
		sentences = append(sentences, sentences2...)
	} else {
		log.Log(log.TopicError, err.Error())
	}

	for iter.Next() == nil {
		ri, err := iter.Open()
		if err != nil {
			return nil, err
		}
		sentences2, err := parseSentences(dict, ri)
		if err == nil {
			sentences = append(sentences, sentences2...)
		} else {
			log.Log(log.TopicError, err.Error())
		}
	}

	return sentences, nil
}
