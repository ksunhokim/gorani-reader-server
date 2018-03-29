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
	r.POST("/wordbooks/:name", view.CreateWordBook)
	r.GET("/wordbooks/:name", view.GetWordBook)
	r.POST("/upload", view.Upload)
	return r
}
