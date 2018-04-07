package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/pkg/router/middlewares"
	"github.com/sunho/engbreaker/pkg/view"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Auth())
	r.GET("/wordbooks", view.ListWordbooks)
	r.POST("/wordbooks/:name", view.CreateWordbook)
	r.GET("/wordbooks/:name", view.GetWordbook)
	r.POST("/wordbooks/:name/words", view.AddEntryToWordbook)
	r.PUT("/wordbooks/:name/words", view.PutEntryToWordbook)
	r.DELETE("/wordbooks/:name", view.DeleteWordbook)
	return r
}
