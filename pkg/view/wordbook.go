package view

import (
	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/pkg/model"
	"github.com/sunho/engbreaker/pkg/router/middlewares"
	"gopkg.in/mgo.v2/bson"
)

func ListWordbooks(c *gin.Context) {
	user := middlewares.User(c)
	c.JSON(200, gin.H{
		"wordbooks": user.Wordbooks,
	})
}

func CreateWordBook(c *gin.Context) {
	name := c.Param("name")
	user := middlewares.User(c)
	book := model.Wordbook{
		Name:    name,
		UserID:  user.GetId(),
		Entries: []model.WordbookEntry{},
	}

	err := model.Save(&book)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	wordbooks := []string{name}
	user.Wordbooks = append(wordbooks, user.Wordbooks...)
	model.Save(&user)
	c.Status(201)
}

func GetWordBook(c *gin.Context) {
	name := c.Param("name")
	book := model.Wordbook{}
	user := middlewares.User(c)
	err := model.Get(&book, bson.M{
		"userid": user.GetId(),
		"name":   name,
	})
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, gin.H{
		"name":     book.Name,
		"entries":  book.Entries,
		"created":  book.GetCreated(),
		"modified": book.GetModified(),
	})
}
