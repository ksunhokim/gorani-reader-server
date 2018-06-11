package book

import (
	"errors"
	"io"
	"strings"

	"github.com/sunho/epubgo"
	"github.com/sunho/gorani-reader-server/go/pkg/log"
	"github.com/sunho/gorani-reader-server/go/pkg/sentencer"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

var (
	ErrNoName        = errors.New("book: Epub doesn't have name")
	ErrMultipleNames = errors.New("book: Epub metadata has multiple names")
)

func Parse(isbn string, r io.ReaderAt, size int64) (b Book, err error) {
	b.Isbn = isbn
	err = b.parseEpub(r, size)
	if err != nil {
		return
	}

	b.fetchInfos()

	return
}

func (b *Book) parseEpub(r io.ReaderAt, size int64) error {
	epub, err := epubgo.Load(r, size)
	if err != nil {
		return err
	}

	err = b.parseMeta(epub)
	if err != nil {
		return err
	}

	err = b.parseSentences(epub)
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) parseMeta(epub *epubgo.Epub) error {
	names, err := epub.Metadata("title")
	if err != nil {
		return ErrNoName
	}

	if len(names) != 1 {
		return ErrMultipleNames
	}

	b.Name = names[0]

	authers, err := epub.Metadata("creator")
	if err == nil {
		b.Author = authers[0]
	}

	cover := epub.Cover()
	if cover != "" {
		name := epub.FileName(cover)
		r, err := epub.OpenFile(name)
		if err != nil {
			return err
		}

		arr := strings.Split(name, ".")
		b.Cover.Ext = arr[len(arr)-1]
		b.Cover.Reader = r
	}

	// TODO use properties

	return nil
}

func (b *Book) parseSentences(epub *epubgo.Epub) error {
	iter, err := epub.Spine()
	if err != nil {
		return err
	}

	for {
		url := iter.URL()
		// xhtml, html
		if !strings.HasSuffix(url, "html") {
			log.Log(log.TopicError, util.M{
				"msg":  "name of file in epub spine doesn't end with html",
				"path": url,
			})
			continue
		}

		r2, err := iter.Open()
		if err != nil {
			return err
		}

		sens, err := sentencer.ExtractSentencesFromHtml(r2)
		if err != nil {
			return err
		}
		b.Sentences = append(b.Sentences, sens...)

		err = iter.Next()
		if err != nil {
			break
		}
	}
	return nil
}
