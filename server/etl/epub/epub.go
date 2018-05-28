package epub

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"strings"
)

type Epub struct {
	files []*zip.File
	index int
}

func New(r io.ReaderAt, size int64) (*Epub, error) {
	z, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}

	e := &Epub{
		files: z.File,
		index: -1,
	}
	return e, nil
}

func (e *Epub) Iterate() bool {
	e.index++
	if e.index == len(e.files) {
		return false
	}

	file := e.files[e.index]
	if !strings.HasSuffix(file.Name, ".html") {
		return e.Iterate()
	}

	return true
}

func (e *Epub) Read() (string, error) {
	file := e.files[e.index]

	r, err := file.Open()
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
