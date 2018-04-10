package view

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/api/model"
	"github.com/sunho/engbreaker/api/router/middlewares"
)

type wordbookListItem struct {
	Name      string     `json:"name"`
	Entries   int        `json:"entries"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ListWordbooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("p"))

	user := middlewares.User(c)
	wordbooks := user.GetWordbooks(page)
	if len(wordbooks) == 0 {
		c.AbortWithStatus(404)
		return
	}

	out := []wordbookListItem{}
	for _, wordbook := range wordbooks {
		out = append(out, wordbookListItem{
			Name:      wordbook.Name,
			Entries:   len(wordbook.Entries),
			UpdatedAt: wordbook.UpdatedAt,
		})
	}

	c.JSON(200, out)
}

func CreateWordbook(c *gin.Context) {
	name := c.Param("name")
	user := middlewares.User(c)

	err := user.CreateWordbook(name)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.Status(201)
}

type entryWithDef struct {
	model.WordbookEntry
	DefinitionText string `json:"definition_text"`
}

func GetWordbook(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))

	user := middlewares.User(c)
	wordbook, err := user.GetWordbook(index)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	entries := []entryWithDef{}
	for _, entry := range wordbook.Entries {
		word, err := model.GetWord(entry.Word)
		if err != nil {
			continue
		}
		entries = append(entries, entryWithDef{
			WordbookEntry:  entry,
			DefinitionText: word.Definitions[entry.Definition].Definition,
		})
	}

	c.JSON(200, gin.H{
		"name":       wordbook.Name,
		"entries":    entries,
		"updated_at": wordbook.UpdatedAt,
	})
}

func PutEntriesOfWordbook(c *gin.Context) {
	index_ := c.Param("index")
	index, _ := strconv.Atoi(index_)

	user := middlewares.User(c)
	wordbook, err := user.GetWordbook(index)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	req := []model.WordbookEntry{}
	err = c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = wordbook.PutEntries(req)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.Status(200)
}

func DeleteWordbook(c *gin.Context) {
	index, _ := strconv.Atoi(c.Param("index"))

	user := middlewares.User(c)
	err := user.DeleteWordbook(index)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Status(200)
}
