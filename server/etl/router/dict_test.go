package router_test

import "testing"

func TestGetWords(t *testing.T) {
	e, s := prepareServer(t)
	defer s.Close()

	resp := e.GET("/word").
		Expect().
		Status(200).
		JSON().
		Array()
	resp.Length().Equal(1)

	word := resp.Element(0).Object()
	word.Keys().ContainsOnly("id", "word")
	word.Value("id").Equal(1)
	word.Value("word").Equal("test")
}

func TestGetWord(t *testing.T) {
	e, s := prepareServer(t)
	defer s.Close()

	resp := e.GET("/word/1").
		Expect().
		Status(200).
		JSON().
		Object()
	resp.Keys().ContainsOnly("id", "word", "definitions")
	defs := resp.Value("definitions").Array()
	defs.Length().Equal(2)

	def := defs.Element(0).Object()
	def.Keys().ContainsOnly("id", "word_id", "definition")
}

func TestDeleteWord(t *testing.T) {
	e, s := prepareServer(t)
	defer s.Close()

	e.DELETE("/word/1").
		Expect().
		Status(200)

	e.GET("/word").
		Expect().
		Status(200).
		JSON().
		Array().
		Empty()

	e.GET("/word/1").
		Expect().
		Status(404)
}

func TestAddWord(t *testing.T) {
	e, s := prepareServer(t)
	defer s.Close()

	req := util.M {
		"word": "hoi!",
		"pronunciation": "fjdfaf",
		"definitions": []util.M {
			util.M{
				"definition": "hai",
				""
			},
		},
	}
	e.POST("/word").

}
