package selector

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	a := assert.New(t)
	s := newScanner(strings.NewReader("ab.cd"))
	p := newParser(s)

	obj, _ := p.parse()
	a.Equal(obj.name, "ab")

	obj, _ = p.parse()
	a.Equal(obj.name, "cd")

	s = newScanner(strings.NewReader("*.cd.*"))
	p = newParser(s)

	obj, _ = p.parse()
	a.True(obj.asterisk)

	obj, _ = p.parse()
	a.Equal(obj.name, "cd")

	obj, _ = p.parse()
	a.True(obj.asterisk)

	s = newScanner(strings.NewReader("ds[10][*]"))
	p = newParser(s)

	obj, _ = p.parse()
	a.Equal(obj.name, "ds")

	index := obj.index
	a.Equal(index.pos, 10)

	index = index.next
	a.True(index.asterisk)

	s = newScanner(strings.NewReader("*.a[*][2].asdf.d[40][*]"))
	p = newParser(s)

	obj, _ = p.parse()
	a.True(obj.asterisk)

	obj, _ = p.parse()
	a.Equal(obj.name, "a")
	a.True(obj.index.asterisk)
	a.Equal(obj.index.next.pos, 2)

	obj, _ = p.parse()
	a.Equal(obj.name, "asdf")

	obj, _ = p.parse()
	a.Equal(obj.name, "d")
	a.Equal(obj.index.pos, 40)
	a.True(obj.index.next.asterisk)
}
