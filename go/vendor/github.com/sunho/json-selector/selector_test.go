package selector_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	selector "github.com/sunho/json-selector"
)

func TestStrings(t *testing.T) {
	a := assert.New(t)
	text := []byte(`
		{
			"hello": "hi",
			"key": "ko",
			"a": {
				"b": "c"
			},
			"한글": "ok"
		}
	`)
	str, err := selector.Select(text, "hello")
	a.Nil(err)
	a.Equal(string(str), "hi")

	str, err = selector.Select(text, "key")
	a.Nil(err)
	a.Equal(string(str), "ko")

	str, err = selector.Select(text, "a.b")
	a.Nil(err)
	a.Equal(string(str), "c")

	str, err = selector.Select(text, "한글")
	a.Nil(err)
	a.Equal(string(str), "ok")
}

func TestObjects(t *testing.T) {
	a := assert.New(t)
	text := []byte(`
		{
			"hello":{
				"a":"b",
				"c":"d"
			},
			"hello2":{
				"a":{
					"ip":"1",
					"pr":"2"
				},
				"b":{
					"ip":"3",
					"pr":"4"
				}
			}
		}
	`)
	hello := struct {
		A string `json:"a"`
		C string `json:"c"`
	}{}

	hello2 := struct {
		A struct {
			Ip string `json:"ip"`
			Pr string `json:"pr"`
		} `json:"a"`
		B struct {
			Ip string `json:"ip"`
			Pr string `json:"pr"`
		} `json:"b"`
	}{}

	str, _ := selector.Select(text, "hello")
	fmt.Println(string(str))
	err := json.Unmarshal(str, &hello)
	a.Nil(err)
	a.Equal(hello.A, "b")
	a.Equal(hello.C, "d")

	str, _ = selector.Select(text, "hello2")
	err = json.Unmarshal(str, &hello2)
	a.Nil(err)
	a.Equal(hello2.A.Ip, "1")
	a.Equal(hello2.A.Pr, "2")
	a.Equal(hello2.B.Ip, "3")
	a.Equal(hello2.B.Pr, "4")
}
