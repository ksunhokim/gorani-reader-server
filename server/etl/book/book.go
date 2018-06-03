package book

import (
	"io"
)

type Book struct {
	Title       string
	Author      string
	Genre       string
	CoverExt    string
	CoverReader io.Reader
	Reviews     []Review
}

type Review struct {
	Provider string
	Number   int
	Rate     float32
}
