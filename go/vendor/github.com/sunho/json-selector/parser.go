package selector

import (
	"fmt"
	"io"
	"strconv"
)

type parser struct {
	s   *scanner
	tok token
	lit string
}

func newParser(s *scanner) *parser {
	return &parser{s: s}
}

type object struct {
	asterisk bool
	name     string
	index    *index
}

type index struct {
	asterisk bool
	pos      int
	next     *index
}

func (p *parser) parse() (object, error) {
	var (
		ltok = tEof
		obj  = object{}
	)

	for {
		p.scan()

		if p.tok == tEof && ltok == tEof {
			return obj, fmt.Errorf("empty name")
		}

		if ltok == tBackslash {
			if p.tok == tBackslash || p.tok == tAsterisk ||
				p.tok == tClose || p.tok == tOpen ||
				p.tok == tDot {
				obj.name += p.lit
				ltok = tLetter
				continue
			}
			return obj, fmt.Errorf("unrecongnized escape")
		}

		if p.tok == tEof {
			return obj, io.EOF
		}

		if p.tok == tAsterisk {
			if (p.peak() == tDot || p.peak() == tEof) && ltok == tEof {
				obj.asterisk = true
				p.scan()
				return obj, nil
			}
			return obj, fmt.Errorf("invalid asterisk")
		}

		if p.tok == tClose {
			return obj, fmt.Errorf("closing before opening")
		}

		if p.tok == tOpen {
			break
		}

		if p.tok == tDot {
			return obj, nil
		}

		obj.name += p.lit
		ltok = p.tok
	}

	opening := true
	lit := ""
	currentIndex := &index{}
	obj.index = currentIndex

	for {
		p.scan()
		if opening {
			if p.tok == tOpen {
				return obj, fmt.Errorf("opening again")
			}

			if p.tok == tAsterisk && p.peak() == tClose && lit == "" {
				currentIndex.asterisk = true

				opening = false
				p.scan()
				continue
			}

			if p.tok == tClose {
				if lit == "" {
					return obj, fmt.Errorf("empty index")
				}

				i, _ := strconv.Atoi(lit)
				currentIndex.pos = i

				lit = ""
				opening = false
				continue
			}

			if p.tok == tNumber {
				lit += p.lit
				continue
			}

		} else {
			if p.tok == tClose {
				return obj, fmt.Errorf("closing again")
			}

			if p.tok == tOpen {
				newIndex := &index{}
				currentIndex.next = newIndex
				currentIndex = newIndex
				opening = true

				continue
			}

			if p.tok == tDot || p.tok == tEof {
				return obj, nil
			}
		}

		return obj, fmt.Errorf("invalid index")
	}
}

func (p *parser) scan() {
	tok, lit := p.s.scan()
	p.tok = tok
	p.lit = lit
}

func (p *parser) peak() token {
	tok, _ := p.s.scan()
	p.s.unread()
	return tok
}
