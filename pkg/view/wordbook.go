package view

import (
	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/pkg/model"
	"github.com/sunho/engbreaker/pkg/router/middlewares"
)

func RetrieveWordBooks(c *gin.Context) {
	type Response struct {
	}
	user := middlewares.User(c)
	c.JSON(200, gin.H{
		"wordbooks": user.Wordbooks,
	})
}

func CreateWordBook(c *gin.Context) {
	name := c.Param("name")
	book := model.Wordbook{
		Name:    name,
		UserID:  middlewares.User(c).GetId(),
		Entries: []model.WordbookEntry{},
	}

	err := model.Save(&book)
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(201)
}
