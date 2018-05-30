package book

import (
	"io"

	"github.com/meskio/epubgo"
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

	bytes, err := getReaderBy(iter)
	if err != nil {
		return nil, err
	}

	for iter.Next() == nil {
		ri, err := iter.Open()
		if err != nil {
			return []io.Reader{}, err
		}

		out = append(out, ri)
	}
	return nil, nil
}

func parseSentences(dict Dictionary, r io.Reader) ([]Sentence, error) {
	out := []io.Reader{ri}
	return out, nil
}
