package view

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/api/model"
	"github.com/sunho/engbreaker/api/router/middlewares"
	"github.com/sunho/engbreaker/pkg/dbs"
)

func ListWordbooks(c *gin.Context) {
	user := middlewares.User(c)
	c.JSON(200, user.Wordbooks)
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
	index_ := c.Param("index")
	index, _ := strconv.Atoi(index_)

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
		"name":     wordbook.Name,
		"entries":  entries,
		"modified": wordbook.UpdatedAt,
	})
}

func PutEntryToWordbook(c *gin.Context) {
	index_ := c.Param("index")
	index, _ := strconv.Atoi(index_)

	user := middlewares.User(c)
	book, err := user.GetWordbook(index)
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

	for _, entry := range req {
		if !model.ValidateWord(entry.WordRef) {
			c.AbortWithStatus(400)
			return
		}
	}

	book.Entries = req
	model.Save(&book)
	c.Status(200)
}

func DeleteWordbook(c *gin.Context) {
	name := c.Param("name")
	user := middlewares.User(c)
	book, err := user.GetWordbook(name)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	err = dbs.MDB.Collection("wordbooks").DeleteDocument(&book)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	model.Save(&user)
	c.Status(200)
}
