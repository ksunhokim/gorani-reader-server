package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/pkg/router/middlewares"
	"github.com/sunho/engbreaker/pkg/view"
)

func New() http.Handler {
	r := gin.New()
	r.Use(middlewares.Auth())
	r.GET("/wordbooks", view.ListWordbooks)
	r.POST("/wordbooks/:name", view.CreateWordBook)
	r.GET("/wordbooks/:name", view.GetWordBook)
	return r
}
