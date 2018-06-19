package selector

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/antonholmquist/jason"
)

func selectOne(obj *jason.Object, p *parser) (interface{}, error) {
	o, err := p.parse()

	if err == io.EOF {
		val, err := obj.GetValue(o.name)
		if err != nil {
			return nil, err
		}
		return val.Interface(), nil
	}

	if err != nil {
		return nil, err
	}

	obj, err = obj.GetObject(o.name)
	if err != nil {
		return nil, err
	}

	i, err := selectOne(obj, p)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func Select(bytes []byte, sel string) ([]byte, error) {
	v, err := jason.NewObjectFromBytes(bytes)
	if err != nil {
		return []byte{}, err
	}

	s := newScanner(strings.NewReader(sel))
	p := newParser(s)

	val, err := selectOne(v, p)
	if err != nil {
		return []byte{}, err
	}

	switch val.(type) {
	case map[string]interface{}:
		out, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		return out, nil
	case []interface{}:
		out, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		return out, nil
	default:
		str := fmt.Sprintf("%v", val)
		return []byte(str), nil
	}
}
