package epub_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/input/epub"
)

func TestEpub(t *testing.T) {
	a := assert.New(t)
	file, err := os.Open("test.epub")
	a.Nil(err)

	info, err := file.Stat()
	a.Nil(err)

	size := info.Size()

	e, err := epub.New(file, size)
	a.Nil(err)

	ok := false
	for e.Iterate() {
		text, err := e.Read()
		a.Nil(err)

		if strings.Contains(text, `<?xml version="1.0" encoding="UTF-8" ?>`) {
			ok = true
		}
	}

	a.Equal(true, ok)
}
