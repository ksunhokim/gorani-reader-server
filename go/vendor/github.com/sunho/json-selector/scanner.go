package selector

import (
	"bufio"
	"io"
)

type token int

const (
	tAsterisk token = iota
	tNumber
	tBackslash
	tOpen
	tClose
	tLetter
	tDot
	tEof
)

type scanner struct {
	r *bufio.Reader
}

func isNumber(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}

func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r)}
}

func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return 0
	}
	return ch
}

func (s *scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *scanner) scan() (token, string) {
	ch := s.read()

	if isNumber(ch) {
		return tNumber, string(ch)
	}

	switch ch {
	case 0:
		return tEof, ""
	case '\\':
		return tBackslash, string(ch)
	case '[':
		return tOpen, string(ch)
	case ']':
		return tClose, string(ch)
	case '.':
		return tDot, string(ch)
	case '*':
		return tAsterisk, string(ch)
	default:
		return tLetter, string(ch)
	}
}
