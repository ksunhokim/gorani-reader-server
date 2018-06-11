package sentencer

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenKind int

const (
	TokenKindEof = iota
	TokenKindEos
	TokenKindCapitalWord
	TokenKindNormalWord
	TokenKindPunc
	TokenKindUnknown
	TokenKindBlank
)

type (
	Token struct {
		Kind TokenKind
		Lit  string
	}

	Tokenizer struct {
		DotSpecialCases []map[string]bool
		s               *bufio.Scanner
		toks            []Token
		// store read data to buffer
		// in order to implement unread
		head   int
		buffer []Token
	}

	// abbrs
	// [index in word array splitted by dot]map[word](true if it is the end of abbr)
	DotSpecialCases []map[string]bool
)

func NewTokenizer(r io.Reader) *Tokenizer {
	s := bufio.NewScanner(r)
	s.Split(split)
	return &Tokenizer{
		DotSpecialCases: DotSpecialCases{},
		s:               s,
		buffer:          []Token{},
		toks:            []Token{},
	}
}

func isWordToken(t Token) bool {
	return t.Kind == TokenKindCapitalWord || t.Kind == TokenKindNormalWord
}

func isStartingQuote(t Token) bool {
	return strings.HasPrefix(t.Lit, "“") || strings.HasPrefix(t.Lit, "\"")
}

func isEndingQuote(t Token) bool {
	return strings.HasSuffix(t.Lit, "”") || strings.HasSuffix(t.Lit, "\"")
}

func (t *Tokenizer) Tokenize() []Token {
	for {
		tok := t.read()
		t.toks = append(t.toks, tok)
		if tok.Kind == TokenKindEof {
			return t.toks
		}

		// this combines word dash word
		// eg) three-dimentional
		if tok.Lit == "-" && isWordToken(t.lastToken(1)) {
			next := t.read()
			if isWordToken(next) {
				t.toks = t.toks[:len(t.toks)-1]
				t.toks[len(t.toks)-1].Lit += tok.Lit + next.Lit
				continue
			}
			t.unread()
		}

		// abbr
		if tok.Lit == "." && isWordToken(t.lastToken(1)) {
			if lit := t.scanDotSpecialCase(); lit != "" {
				t.toks = t.toks[:len(t.toks)-1]
				t.toks[len(t.toks)-1].Lit = lit
				continue
			}
		}

		// if the next token is capitalized word, " ends the sentence
		if isEndingQuote(tok) {
			if t.scanQuoteEos() {
				t.toks = append(t.toks, Token{TokenKindEos, ""})
				continue
			}
		}

		// if it is not abbr, dot ends the sentence
		if strings.HasSuffix(tok.Lit, ".") || strings.HasSuffix(tok.Lit, "?") || strings.HasSuffix(tok.Lit, "!") {
			t.toks = append(t.toks, Token{TokenKindEos, ""})
			continue
		}
	}
}

// lastToken(0) -> current token
// lastToken(n) -> n previous token
// returns Eof if n is invalid
func (t *Tokenizer) lastToken(n int) Token {
	if n >= len(t.toks) {
		return Token{TokenKindEof, ""}
	}

	return t.toks[len(t.toks)-n-1]
}

func (t *Tokenizer) hasDotSpecialCase(index int, word string) (bool, bool) {
	if index >= len(t.DotSpecialCases) {
		return false, false
	}

	isend, ok := t.DotSpecialCases[index][strings.ToLower(word)]
	return isend, ok
}

func (t *Tokenizer) scanDotSpecialCase() (lit string) {
	var (
		tok = t.lastToken(1)
		// index in DotSpecialCases
		index = 1
		// candidate lit that might be part of complete abbr
		clit string
		// in order to roll back the reads
		// that happened to check if clit is part of complete abbr
		readn int
		// fsm state
		// 0: reads blank
		// 1: reads word
		// 2: reads dot
		state int
	)

	if isend, ok := t.hasDotSpecialCase(0, tok.Lit); !ok {
		return
	} else {
		if isend {
			lit = tok.Lit + "."
		} else {
			clit = tok.Lit + "."
		}
	}

	for {
		tok = t.read()
		readn++
		if state == 0 {
			state = 1
			if isWordToken(tok) {
				// accept the abbr that doesn't have
				// spaces between them
				t.unread()
				readn--
			} else if tok.Kind != TokenKindBlank {
				break
			}
		} else if state == 1 {
			if isWordToken(tok) {
				if isend, ok := t.hasDotSpecialCase(index, tok.Lit); ok {
					index++
					state = 2

					if isend {
						lit += clit + tok.Lit
						clit = ""
						// set to 0 to not roll back because
						// read tokens were already processed
						readn = 0
					} else {
						clit += tok.Lit
					}
					continue
				}
			}
			break
		} else if state == 2 {
			if tok.Lit == "." {
				state = 0
				// clit is empty if the prior state 1
				// determined a complete abbr
				if clit == "" {
					lit += "."
					// set to 0 to not roll back because
					// read tokens were already processed
					readn = 0
				} else {
					clit += "."
				}
			} else {
				break
			}
		}
	}

	// roll back
	for i := 0; i < readn; i++ {
		t.unread()
	}
	return
}

func (t *Tokenizer) scanQuoteEos() bool {
	// check if it is starting quotes
	isStarting := true
	for i := len(t.toks) - 2; i >= 0; i-- {
		tok := t.toks[i]
		if isStartingQuote(tok) {
			isStarting = false
			break
		}

		if tok.Kind == TokenKindEos {
			break
		}
	}

	if isStarting {
		return false
	}

	for {
		tok := t.read()
		defer t.unread()

		switch tok.Kind {
		case TokenKindBlank:
			continue
		case TokenKindCapitalWord:
			return true
		case TokenKindPunc:
			if isStartingQuote(tok) {
				return true
			}
			return false
		default:
			return false
		}
	}
}

func (t *Tokenizer) unread() {
	if t.head != 0 {
		t.head--
	}
}

func (t *Tokenizer) read() Token {
	if t.head == len(t.buffer) {
		// load more data to buffer
		var tok Token
		if t.s.Scan() {
			s := t.s.Bytes()
			u, _ := utf8.DecodeRune(s)
			tok.Lit = string(s)

			if unicode.IsSpace(u) {
				tok.Kind = TokenKindBlank
			} else if unicode.IsPunct(u) {
				tok.Kind = TokenKindPunc
			} else if unicode.IsUpper(u) {
				tok.Kind = TokenKindCapitalWord
			} else if unicode.IsLower(u) {
				tok.Kind = TokenKindNormalWord
			} else {
				tok.Kind = TokenKindUnknown
			}
		} else {
			tok = Token{TokenKindEof, ""}
		}
		t.buffer = append(t.buffer, tok)
	}

	tok := t.buffer[t.head]
	t.head++
	return tok
}

func isAlphabet(u rune) bool {
	return 'a' <= u && u <= 'z' || 'A' <= u && u <= 'Z'
}

func isUnknown(u rune) bool {
	return !unicode.IsPunct(u) && !isAlphabet(u) && !unicode.IsSpace(u)
}

func splitOne(eof bool, data []byte, start int, u rune, shouldAdvance func(u rune) bool) (int, []byte, error) {
	for width, i := 0, start; i < len(data) && shouldAdvance(u); i += width {
		u, width = utf8.DecodeRune(data[i:])
		if !shouldAdvance(u) {
			return i, data[:i], nil
		}
	}

	if eof {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func split(data []byte, atEOF bool) (advance int, Token []byte, err error) {
	u, start := utf8.DecodeRune(data)
	if unicode.IsSpace(u) {
		return splitOne(atEOF, data, start, u, unicode.IsSpace)
	}

	if unicode.IsPunct(u) {
		return splitOne(atEOF, data, start, u, unicode.IsPunct)
	}

	if isAlphabet(u) {
		return splitOne(atEOF, data, start, u, isAlphabet)
	}

	if isUnknown(u) && start != 0 {
		return splitOne(atEOF, data, start, u, isUnknown)
	}

	return
}
